// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: region.proto

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

// RegionServiceClient is the client API for RegionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegionServiceClient interface {
	CreateRegion(ctx context.Context, in *CreateRegionRequest, opts ...grpc.CallOption) (*CreateRegionResponse, error)
	GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetRegionResponse, error)
	DeleteRegion(ctx context.Context, in *DeleteRegionRequest, opts ...grpc.CallOption) (*DeleteRegionResponse, error)
	ListRegion(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*ListRegionsResponse, error)
}

type regionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionServiceClient(cc grpc.ClientConnInterface) RegionServiceClient {
	return &regionServiceClient{cc}
}

func (c *regionServiceClient) CreateRegion(ctx context.Context, in *CreateRegionRequest, opts ...grpc.CallOption) (*CreateRegionResponse, error) {
	out := new(CreateRegionResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.RegionService/CreateRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetRegionResponse, error) {
	out := new(GetRegionResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.RegionService/GetRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) DeleteRegion(ctx context.Context, in *DeleteRegionRequest, opts ...grpc.CallOption) (*DeleteRegionResponse, error) {
	out := new(DeleteRegionResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.RegionService/DeleteRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) ListRegion(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*ListRegionsResponse, error) {
	out := new(ListRegionsResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.RegionService/ListRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionServiceServer is the server API for RegionService service.
// All implementations must embed UnimplementedRegionServiceServer
// for forward compatibility
type RegionServiceServer interface {
	CreateRegion(context.Context, *CreateRegionRequest) (*CreateRegionResponse, error)
	GetRegion(context.Context, *GetRegionRequest) (*GetRegionResponse, error)
	DeleteRegion(context.Context, *DeleteRegionRequest) (*DeleteRegionResponse, error)
	ListRegion(context.Context, *ListRegionsRequest) (*ListRegionsResponse, error)
	mustEmbedUnimplementedRegionServiceServer()
}

// UnimplementedRegionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegionServiceServer struct {
}

func (UnimplementedRegionServiceServer) CreateRegion(context.Context, *CreateRegionRequest) (*CreateRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegion not implemented")
}
func (UnimplementedRegionServiceServer) GetRegion(context.Context, *GetRegionRequest) (*GetRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegion not implemented")
}
func (UnimplementedRegionServiceServer) DeleteRegion(context.Context, *DeleteRegionRequest) (*DeleteRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegion not implemented")
}
func (UnimplementedRegionServiceServer) ListRegion(context.Context, *ListRegionsRequest) (*ListRegionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegion not implemented")
}
func (UnimplementedRegionServiceServer) mustEmbedUnimplementedRegionServiceServer() {}

// UnsafeRegionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegionServiceServer will
// result in compilation errors.
type UnsafeRegionServiceServer interface {
	mustEmbedUnimplementedRegionServiceServer()
}

func RegisterRegionServiceServer(s grpc.ServiceRegistrar, srv RegionServiceServer) {
	s.RegisterService(&RegionService_ServiceDesc, srv)
}

func _RegionService_CreateRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).CreateRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.RegionService/CreateRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).CreateRegion(ctx, req.(*CreateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_GetRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).GetRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.RegionService/GetRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).GetRegion(ctx, req.(*GetRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_DeleteRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).DeleteRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.RegionService/DeleteRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).DeleteRegion(ctx, req.(*DeleteRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_ListRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).ListRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.RegionService/ListRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).ListRegion(ctx, req.(*ListRegionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegionService_ServiceDesc is the grpc.ServiceDesc for RegionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "online_shop.management.v1.RegionService",
	HandlerType: (*RegionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRegion",
			Handler:    _RegionService_CreateRegion_Handler,
		},
		{
			MethodName: "GetRegion",
			Handler:    _RegionService_GetRegion_Handler,
		},
		{
			MethodName: "DeleteRegion",
			Handler:    _RegionService_DeleteRegion_Handler,
		},
		{
			MethodName: "ListRegion",
			Handler:    _RegionService_ListRegion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "region.proto",
}
