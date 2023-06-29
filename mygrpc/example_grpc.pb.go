// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: example.proto

package mygrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RoiuteMessage_SayHello_FullMethodName = "/RoiuteMessage/SayHello"
)

// RoiuteMessageClient is the client API for RoiuteMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoiuteMessageClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type roiuteMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewRoiuteMessageClient(cc grpc.ClientConnInterface) RoiuteMessageClient {
	return &roiuteMessageClient{cc}
}

func (c *roiuteMessageClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, RoiuteMessage_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoiuteMessageServer is the server API for RoiuteMessage service.
// All implementations must embed UnimplementedRoiuteMessageServer
// for forward compatibility
type RoiuteMessageServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedRoiuteMessageServer()
}

// UnimplementedRoiuteMessageServer must be embedded to have forward compatible implementations.
type UnimplementedRoiuteMessageServer struct {
}

func (UnimplementedRoiuteMessageServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedRoiuteMessageServer) mustEmbedUnimplementedRoiuteMessageServer() {}

// UnsafeRoiuteMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoiuteMessageServer will
// result in compilation errors.
type UnsafeRoiuteMessageServer interface {
	mustEmbedUnimplementedRoiuteMessageServer()
}

func RegisterRoiuteMessageServer(s grpc.ServiceRegistrar, srv RoiuteMessageServer) {
	s.RegisterService(&RoiuteMessage_ServiceDesc, srv)
}

func _RoiuteMessage_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoiuteMessageServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoiuteMessage_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoiuteMessageServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoiuteMessage_ServiceDesc is the grpc.ServiceDesc for RoiuteMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoiuteMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RoiuteMessage",
	HandlerType: (*RoiuteMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _RoiuteMessage_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}
