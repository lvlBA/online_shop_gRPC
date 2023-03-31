// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: location.proto

package management

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

// LocationServiceClient is the client API for LocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationServiceClient interface {
	CreateLocation(ctx context.Context, in *CreateLocationRequest, opts ...grpc.CallOption) (*CreateLocationResponse, error)
	GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error)
	DeleteLocation(ctx context.Context, in *DeleteLocationRequest, opts ...grpc.CallOption) (*DeleteLocationResponse, error)
	ListLocation(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error)
}

type locationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationServiceClient(cc grpc.ClientConnInterface) LocationServiceClient {
	return &locationServiceClient{cc}
}

func (c *locationServiceClient) CreateLocation(ctx context.Context, in *CreateLocationRequest, opts ...grpc.CallOption) (*CreateLocationResponse, error) {
	out := new(CreateLocationResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.LocationService/CreateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error) {
	out := new(GetLocationResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.LocationService/GetLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) DeleteLocation(ctx context.Context, in *DeleteLocationRequest, opts ...grpc.CallOption) (*DeleteLocationResponse, error) {
	out := new(DeleteLocationResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.LocationService/DeleteLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) ListLocation(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error) {
	out := new(ListLocationsResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.LocationService/ListLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationServiceServer is the server API for LocationService service.
// All implementations must embed UnimplementedLocationServiceServer
// for forward compatibility
type LocationServiceServer interface {
	CreateLocation(context.Context, *CreateLocationRequest) (*CreateLocationResponse, error)
	GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error)
	DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error)
	ListLocation(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error)
	mustEmbedUnimplementedLocationServiceServer()
}

// UnimplementedLocationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocationServiceServer struct {
}

func (UnimplementedLocationServiceServer) CreateLocation(context.Context, *CreateLocationRequest) (*CreateLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocation not implemented")
}
func (UnimplementedLocationServiceServer) GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedLocationServiceServer) DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocation not implemented")
}
func (UnimplementedLocationServiceServer) ListLocation(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLocation not implemented")
}
func (UnimplementedLocationServiceServer) mustEmbedUnimplementedLocationServiceServer() {}

// UnsafeLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationServiceServer will
// result in compilation errors.
type UnsafeLocationServiceServer interface {
	mustEmbedUnimplementedLocationServiceServer()
}

func RegisterLocationServiceServer(s grpc.ServiceRegistrar, srv LocationServiceServer) {
	s.RegisterService(&LocationService_ServiceDesc, srv)
}

func _LocationService_CreateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).CreateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.LocationService/CreateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).CreateLocation(ctx, req.(*CreateLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.LocationService/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).GetLocation(ctx, req.(*GetLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_DeleteLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).DeleteLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.LocationService/DeleteLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).DeleteLocation(ctx, req.(*DeleteLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_ListLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).ListLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.LocationService/ListLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).ListLocation(ctx, req.(*ListLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationService_ServiceDesc is the grpc.ServiceDesc for LocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "online_shop.management.v1.LocationService",
	HandlerType: (*LocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLocation",
			Handler:    _LocationService_CreateLocation_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _LocationService_GetLocation_Handler,
		},
		{
			MethodName: "DeleteLocation",
			Handler:    _LocationService_DeleteLocation_Handler,
		},
		{
			MethodName: "ListLocation",
			Handler:    _LocationService_ListLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "location.proto",
}