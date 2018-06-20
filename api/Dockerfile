FROM golang:alpine AS build-env
WORKDIR /go/src/github.com/keywee/api
COPY . .
RUN go build -o api

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/keywee/api/api /app/

ENV KAFKA_ADDRESS=kafka:9092
ENV KAFKA_TOPIC=keywee
ENV ELASTIC_ADDRESS=elasticsearch:9200
ENV ELASTIC_INDEX=keywee
ENV PORT=9090

ENTRYPOINT ./api
