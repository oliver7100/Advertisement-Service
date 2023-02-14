// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/service.proto

package proto

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

// AdvertisementServiceClient is the client API for AdvertisementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdvertisementServiceClient interface {
	CreateAdvertisement(ctx context.Context, in *Advertisement, opts ...grpc.CallOption) (*Advertisement, error)
}

type advertisementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdvertisementServiceClient(cc grpc.ClientConnInterface) AdvertisementServiceClient {
	return &advertisementServiceClient{cc}
}

func (c *advertisementServiceClient) CreateAdvertisement(ctx context.Context, in *Advertisement, opts ...grpc.CallOption) (*Advertisement, error) {
	out := new(Advertisement)
	err := c.cc.Invoke(ctx, "/proto.AdvertisementService/CreateAdvertisement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvertisementServiceServer is the server API for AdvertisementService service.
// All implementations must embed UnimplementedAdvertisementServiceServer
// for forward compatibility
type AdvertisementServiceServer interface {
	CreateAdvertisement(context.Context, *Advertisement) (*Advertisement, error)
	mustEmbedUnimplementedAdvertisementServiceServer()
}

// UnimplementedAdvertisementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdvertisementServiceServer struct {
}

func (UnimplementedAdvertisementServiceServer) CreateAdvertisement(context.Context, *Advertisement) (*Advertisement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdvertisement not implemented")
}
func (UnimplementedAdvertisementServiceServer) mustEmbedUnimplementedAdvertisementServiceServer() {}

// UnsafeAdvertisementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdvertisementServiceServer will
// result in compilation errors.
type UnsafeAdvertisementServiceServer interface {
	mustEmbedUnimplementedAdvertisementServiceServer()
}

func RegisterAdvertisementServiceServer(s grpc.ServiceRegistrar, srv AdvertisementServiceServer) {
	s.RegisterService(&AdvertisementService_ServiceDesc, srv)
}

func _AdvertisementService_CreateAdvertisement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Advertisement)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServiceServer).CreateAdvertisement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AdvertisementService/CreateAdvertisement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServiceServer).CreateAdvertisement(ctx, req.(*Advertisement))
	}
	return interceptor(ctx, in, info, handler)
}

// AdvertisementService_ServiceDesc is the grpc.ServiceDesc for AdvertisementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdvertisementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AdvertisementService",
	HandlerType: (*AdvertisementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAdvertisement",
			Handler:    _AdvertisementService_CreateAdvertisement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
