// Code generated by protoc-gen-go.
// source: iot.proto
// DO NOT EDIT!

/*
Package iot is a generated protocol buffer package.

It is generated from these files:
	iot.proto

It has these top-level messages:
	Config
	Response
*/
package iot

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

type Config struct {
	Beat   string `protobuf:"bytes,1,opt,name=beat" json:"beat,omitempty"`
	Action string `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetBeat() string {
	if m != nil {
		return m.Beat
	}
	return ""
}

func (m *Config) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Config) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Response struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Action  string `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	Code    string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	Result  string `protobuf:"bytes,4,opt,name=result" json:"result,omitempty"`
	Message string `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Response) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Response) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "iot.Config")
	proto.RegisterType((*Response)(nil), "iot.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IOT service

type IOTClient interface {
	Beat(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Response, error)
}

type iOTClient struct {
	cc *grpc.ClientConn
}

func NewIOTClient(cc *grpc.ClientConn) IOTClient {
	return &iOTClient{cc}
}

func (c *iOTClient) Beat(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/iot.IOT/Beat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IOT service

type IOTServer interface {
	Beat(context.Context, *Config) (*Response, error)
}

func RegisterIOTServer(s *grpc.Server, srv IOTServer) {
	s.RegisterService(&_IOT_serviceDesc, srv)
}

func _IOT_Beat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IOTServer).Beat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iot.IOT/Beat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IOTServer).Beat(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

var _IOT_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iot.IOT",
	HandlerType: (*IOTServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Beat",
			Handler:    _IOT_Beat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iot.proto",
}

func init() { proto.RegisterFile("iot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0xcc, 0x2f, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0xcc, 0x2f, 0x51, 0xf2, 0xe0, 0x62, 0x73, 0xce,
	0xcf, 0x4b, 0xcb, 0x4c, 0x17, 0x12, 0xe2, 0x62, 0x49, 0x4a, 0x4d, 0x2c, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xc4, 0xb8, 0xd8, 0x12, 0x93, 0x4b, 0x32, 0xf3, 0xf3, 0x24,
	0x98, 0xc0, 0xa2, 0x50, 0x1e, 0x48, 0x6d, 0x4a, 0x62, 0x49, 0xa2, 0x04, 0xb3, 0x02, 0xa3, 0x06,
	0x4f, 0x10, 0x98, 0xad, 0x54, 0xc3, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x0a,
	0x92, 0xcf, 0x4b, 0xcc, 0x4d, 0x85, 0x99, 0x05, 0x62, 0xe3, 0x33, 0x2b, 0x39, 0x3f, 0x25, 0x15,
	0x6c, 0x16, 0x67, 0x10, 0x98, 0x0d, 0x52, 0x5b, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x22, 0xc1, 0x02,
	0x51, 0x0b, 0xe1, 0x09, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0xb0,
	0x82, 0x25, 0x60, 0x5c, 0x23, 0x6d, 0x2e, 0x66, 0x4f, 0xff, 0x10, 0x21, 0x15, 0x2e, 0x16, 0x27,
	0x90, 0xc3, 0xb9, 0xf5, 0x40, 0xfe, 0x84, 0xf8, 0x4c, 0x8a, 0x17, 0xcc, 0x81, 0x39, 0x4e, 0x89,
	0x21, 0x89, 0x0d, 0x1c, 0x00, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x7e, 0x1c, 0x11,
	0x0d, 0x01, 0x00, 0x00,
}
