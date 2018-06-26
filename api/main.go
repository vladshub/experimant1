package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/keywee/api/pb"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"github.com/Shopify/sarama"

	"time"
	"log"
	"strconv"
)

var (
	port *int

	elasticAddress = flag.String("elastic_address", os.Getenv("ELASTIC_ADDRESS"), "ElasticSearch URL")
	index          = flag.String("elastic_index", os.Getenv("ELASTIC_INDEX"), "ElasticSearch index to use")

	logrusLogger = logrus.New()
	logrusEntry  = logrus.NewEntry(logrusLogger)
	// Shared options for the logger, with a custom duration to log field function.
	lopts = []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}
	opts = []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(logrusEntry, lopts...),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, lopts...),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}
	kafkaAddress = flag.String("kafka_address", os.Getenv("KAFKA_ADDRESS"), "Kafka address")
	kafkaTopic = flag.String("kafka_topic", os.Getenv("KAFKA_TOPIC"), "Kafka topic")
	mapping = `
{
	"mappings":{
		"doc":{
			"properties":{
				"id":{
					"type":"text"
				},
				"url":{
					"type":"text"
				},
				"content":{
					"type":"text"
				},
				"introduction":{
					"type":"text"
				},
				"title":{
					"type":"text"
				},
				"fbIntrests":{
					"properties":{
						"path":{
							"type":"text"
						},
						"name":{
							"type":"text"
						},
						"topic":{
							"type":"text"
						},
						"id":{
							"type":"text"
						}
					}
				}
			}
		}
	}
}
`
)

func init() {
	p, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		p = 9090
	}
	port  = flag.Int("port",  p, "The server port")

}

type api struct {
	esClient *elastic.Client
	kafkaProducer sarama.AsyncProducer
	esOptions []elastic.ClientOptionFunc
}

func (s *api) elasticClient() *elastic.Client {
	var err error
	var retries = 0;
BEGINING:
	if s.esClient == nil {
		s.esClient, err = elastic.NewClient(elastic.SetURL(*elasticAddress))
		if err != nil {
			panic(err)
		}
	} else {
		ctx := context.Background()
		info, code, err := s.esClient.Ping(*elasticAddress).Do(ctx)
		if err != nil {
			s.esClient = nil
			if retries < 3 {
				retries++
				goto BEGINING
			} else {
				panic(fmt.Sprintf("Elastic Search Not Available on address %s", *elasticAddress))
			}
		}
		grpclog.Infof("ping code: %d info: %s ", code, info)
	}
	return s.esClient
}

func (s *api) verifyIndex() error {
	exists, err := s.elasticClient().IndexExists(*index).Do(context.Background())
	if err != nil {
		return err
	} else if !exists {
		createIndex, err := s.elasticClient().CreateIndex(*index).BodyString(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			return err
		}
		if !createIndex.Acknowledged {
			return fmt.Errorf("Index creation not acknowledged")
		}
	}
	return nil
}

func (s *api) Add(ctx context.Context, item *pb.Item) (*pb.Status, error) {
	id := uuid.New()
	item.Id = id.String()
	ma := jsonpb.Marshaler{}
	doc, err := ma.MarshalToString(item)
	if err != nil {
		return nil, err
	}

	s.kafkaProducer.Input() <- &sarama.ProducerMessage{
		Topic: *kafkaTopic,
		Key:   sarama.StringEncoder(id.String()),
		Value: sarama.ByteEncoder(doc),
	}

	_, err = s.elasticClient().Index().
		Index(*index).
		Type("doc").
		Id(item.Id).
		BodyJson(doc).
		Refresh("wait_for").
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &pb.Status{
		Code:    200,
		Message: "Added Successfully",
	}, nil
}

func (s *api) Read(ctx context.Context, item *pb.Item) (*pb.Item, error) {
	//var siteType pb.Item
	searchResult, err := s.elasticClient().Get().Index(*index).Id(item.Id).Do(ctx)
	if err != nil {
		return nil, err
	}
	if searchResult.Found {
		data, err := searchResult.Source.MarshalJSON()
		if err != nil {
			return nil, err
		}
		if err := jsonpb.UnmarshalString(string(data), item); err != nil {
			return nil, err
		}
	} else {
		return nil, status.Errorf(codes.NotFound, "Id (%s) was not found", item.Id)
	}
	return item, nil
}

func (s *api) Index(_ *pb.Empty, stream pb.API_IndexServer) error {
	searchResult, err := s.elasticClient().Search().Index(*index).Query(elastic.NewMatchAllQuery()).Do(context.Background())
	if err != nil {
		return err
	}
	stream.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	errors := make([]error, 0)
	if searchResult.TotalHits() > 0 {
		for _, item := range searchResult.Hits.Hits {
			siteType := &pb.Item{}
			data, err := item.Source.MarshalJSON()
			if err != nil {
				errors = append(errors, err)
			}
			if err = jsonpb.UnmarshalString(string(data), siteType); err != nil {
				errors = append(errors, err)
			}
			stream.Send(siteType)
		}
	}

	stream.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

func newServer(kp sarama.AsyncProducer, esOptions ...elastic.ClientOptionFunc) (*api, error) {
	a := &api{
		esOptions:     esOptions,
		kafkaProducer: kp,
	}
	err := a.verifyIndex()
	return a, err
}

func main() {
	flag.Parse()
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	grpcServer := grpc.NewServer(opts...)

	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.ClientID = "keywee-api"

	producer, err := sarama.NewAsyncProducer([]string{*kafkaAddress}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	s, err := newServer(producer, elastic.SetURL(*elasticAddress))
	if err != nil {
		panic(err)
	}
	pb.RegisterAPIServer(grpcServer, s)
	wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithWebsockets(true))
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
		Handler: http.HandlerFunc(handler),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}
