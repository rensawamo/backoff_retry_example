// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/some_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SomeService_SomeMethod_FullMethodName = "/some_service.SomeService/SomeMethod"
)

// SomeServiceClient is the client API for SomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// サービス定義
type SomeServiceClient interface {
	SomeMethod(ctx context.Context, in *SomeRequest, opts ...grpc.CallOption) (*SomeResponse, error)
}

type someServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSomeServiceClient(cc grpc.ClientConnInterface) SomeServiceClient {
	return &someServiceClient{cc}
}

func (c *someServiceClient) SomeMethod(ctx context.Context, in *SomeRequest, opts ...grpc.CallOption) (*SomeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SomeResponse)
	err := c.cc.Invoke(ctx, SomeService_SomeMethod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SomeServiceServer is the server API for SomeService service.
// All implementations must embed UnimplementedSomeServiceServer
// for forward compatibility.
//
// サービス定義
type SomeServiceServer interface {
	SomeMethod(context.Context, *SomeRequest) (*SomeResponse, error)
	mustEmbedUnimplementedSomeServiceServer()
}

// UnimplementedSomeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSomeServiceServer struct{}

func (UnimplementedSomeServiceServer) SomeMethod(context.Context, *SomeRequest) (*SomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SomeMethod not implemented")
}
func (UnimplementedSomeServiceServer) mustEmbedUnimplementedSomeServiceServer() {}
func (UnimplementedSomeServiceServer) testEmbeddedByValue()                     {}

// UnsafeSomeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SomeServiceServer will
// result in compilation errors.
type UnsafeSomeServiceServer interface {
	mustEmbedUnimplementedSomeServiceServer()
}

func RegisterSomeServiceServer(s grpc.ServiceRegistrar, srv SomeServiceServer) {
	// If the following call pancis, it indicates UnimplementedSomeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SomeService_ServiceDesc, srv)
}

func _SomeService_SomeMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SomeServiceServer).SomeMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SomeService_SomeMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SomeServiceServer).SomeMethod(ctx, req.(*SomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SomeService_ServiceDesc is the grpc.ServiceDesc for SomeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "some_service.SomeService",
	HandlerType: (*SomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SomeMethod",
			Handler:    _SomeService_SomeMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/some_service.proto",
}