// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keywee.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Item struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	EstimateReady        bool     `protobuf:"varint,3,opt,name=estimate_ready,json=estimateReady,proto3" json:"estimate_ready,omitempty"`
	AudianceSize         int64    `protobuf:"varint,4,opt,name=audiance_size,json=audianceSize,proto3" json:"audiance_size,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Title                string   `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Introduction         string   `protobuf:"bytes,7,opt,name=introduction,proto3" json:"introduction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_keywee_d6638f78f3f07554, []int{0}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (dst *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(dst, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Item) GetEstimateReady() bool {
	if m != nil {
		return m.EstimateReady
	}
	return false
}

func (m *Item) GetAudianceSize() int64 {
	if m != nil {
		return m.AudianceSize
	}
	return 0
}

func (m *Item) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Item) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Item) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

type Status struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_keywee_d6638f78f3f07554, []int{1}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_keywee_d6638f78f3f07554, []int{2}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Item)(nil), "pb.Item")
	proto.RegisterType((*Status)(nil), "pb.Status")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	Add(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Status, error)
	Read(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error)
	Index(ctx context.Context, in *Empty, opts ...grpc.CallOption) (API_IndexClient, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Add(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/pb.API/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Read(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/pb.API/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Index(ctx context.Context, in *Empty, opts ...grpc.CallOption) (API_IndexClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[0], "/pb.API/Index", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIIndexClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_IndexClient interface {
	Recv() (*Item, error)
	grpc.ClientStream
}

type aPIIndexClient struct {
	grpc.ClientStream
}

func (x *aPIIndexClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	Add(context.Context, *Item) (*Status, error)
	Read(context.Context, *Item) (*Item, error)
	Index(*Empty, API_IndexServer) error
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Add(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Read(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Index_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).Index(m, &aPIIndexServer{stream})
}

type API_IndexServer interface {
	Send(*Item) error
	grpc.ServerStream
}

type aPIIndexServer struct {
	grpc.ServerStream
}

func (x *aPIIndexServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _API_Add_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _API_Read_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Index",
			Handler:       _API_Index_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "keywee.proto",
}

func init() { proto.RegisterFile("keywee.proto", fileDescriptor_keywee_d6638f78f3f07554) }

var fileDescriptor_keywee_d6638f78f3f07554 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0x7f, 0xfd, 0x33, 0xb4, 0x45, 0x06, 0x0f, 0x8b, 0x28, 0x84, 0x15, 0xa1, 0xa7,
	0x22, 0x0a, 0xde, 0x7b, 0xf0, 0xd0, 0x9b, 0xa4, 0x1f, 0xa0, 0x6c, 0xb3, 0x83, 0x2c, 0x36, 0xbb,
	0x21, 0x99, 0xa0, 0xe9, 0x27, 0xf4, 0x63, 0xc9, 0x6e, 0x8d, 0xd4, 0xdb, 0x7b, 0xbf, 0x07, 0x33,
	0xbc, 0x07, 0xf3, 0x0f, 0xea, 0x3f, 0x89, 0xd6, 0x75, 0xe3, 0xd8, 0x61, 0x5c, 0x1f, 0xe4, 0x77,
	0x04, 0xe9, 0x96, 0xa9, 0xc2, 0x25, 0xc4, 0x46, 0x8b, 0x28, 0x8f, 0x56, 0xb3, 0x22, 0x36, 0x1a,
	0xaf, 0x20, 0xe9, 0x9a, 0xa3, 0x88, 0x03, 0xf0, 0x12, 0x1f, 0x60, 0x49, 0x2d, 0x9b, 0x4a, 0x31,
	0xed, 0x1b, 0x52, 0xba, 0x17, 0x49, 0x1e, 0xad, 0xa6, 0xc5, 0x62, 0xa0, 0x85, 0x87, 0x78, 0x0f,
	0x0b, 0xd5, 0x69, 0xa3, 0x6c, 0x49, 0xfb, 0xd6, 0x9c, 0x48, 0xa4, 0x79, 0xb4, 0x4a, 0x8a, 0xf9,
	0x00, 0x77, 0xe6, 0x44, 0x28, 0x60, 0x52, 0x3a, 0xcb, 0x64, 0x59, 0x64, 0xe1, 0xc3, 0x60, 0xf1,
	0x1a, 0x32, 0x36, 0x7c, 0x24, 0x31, 0x0e, 0xfc, 0x6c, 0x50, 0xc2, 0xdc, 0x58, 0x6e, 0x9c, 0xee,
	0x4a, 0x36, 0xce, 0x8a, 0x49, 0x08, 0xff, 0x31, 0xf9, 0x02, 0xe3, 0x1d, 0x2b, 0xee, 0x5a, 0x44,
	0x48, 0x4b, 0xa7, 0x29, 0xb4, 0x49, 0x8a, 0xa0, 0xfd, 0xc7, 0x8a, 0xda, 0x56, 0xbd, 0xd3, 0x6f,
	0xa7, 0xc1, 0xca, 0x09, 0x64, 0xaf, 0x55, 0xcd, 0xfd, 0x93, 0x86, 0x64, 0xf3, 0xb6, 0xc5, 0x3b,
	0x48, 0x36, 0x5a, 0xe3, 0x74, 0x5d, 0x1f, 0xd6, 0x7e, 0x9a, 0x1b, 0xf0, 0xea, 0x7c, 0x5a, 0x8e,
	0xf0, 0x16, 0x52, 0x5f, 0xf4, 0x22, 0xff, 0x53, 0x72, 0x84, 0x39, 0x64, 0x5b, 0xab, 0xe9, 0x0b,
	0x67, 0x1e, 0x86, 0xbb, 0x97, 0xf9, 0x63, 0x74, 0x18, 0x87, 0xf1, 0x9f, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x23, 0x87, 0x42, 0xd0, 0x8c, 0x01, 0x00, 0x00,
}