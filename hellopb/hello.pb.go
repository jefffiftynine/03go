// Code generated by protoc-gen-go.
// source: hello.proto
// DO NOT EDIT!

/*
Package chapter10 is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package chapter10

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
const _ = proto.ProtoPackageIsVersion1

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HelloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "chapter10.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "chapter10.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Hello service

type HelloClient interface {
	Say(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Say(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/chapter10.Hello/Say", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	Say(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(HelloServer).Say(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chapter10.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Hello_Say_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xce, 0x48, 0x2c, 0x28, 0x49, 0x2d,
	0x32, 0x34, 0x50, 0x52, 0xe2, 0xe2, 0xf1, 0x00, 0xc9, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81,
	0xd9, 0x4a, 0x9a, 0x5c, 0xbc, 0x50, 0x35, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12, 0x5c,
	0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x75, 0x30, 0xae, 0x91, 0x33, 0x17, 0x2b, 0x58,
	0xa9, 0x90, 0x15, 0x17, 0x73, 0x70, 0x62, 0xa5, 0x90, 0xb8, 0x1e, 0xdc, 0x2a, 0x3d, 0x64, 0x7b,
	0xa4, 0x24, 0x30, 0x25, 0x20, 0x86, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0x5d, 0x69, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x07, 0xf3, 0x53, 0x0a, 0xb4, 0x00, 0x00, 0x00,
}