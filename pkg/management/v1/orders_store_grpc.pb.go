// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: orders_store.proto

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

// OrdersStoreServiceClient is the client API for OrdersStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersStoreServiceClient interface {
	CreateOrdersStore(ctx context.Context, in *CreateOrdersStoreRequest, opts ...grpc.CallOption) (*CreateOrdersStoreResponse, error)
	GetOrdersStore(ctx context.Context, in *GetOrdersStoreRequest, opts ...grpc.CallOption) (*GetOrdersStoreResponse, error)
	DeleteOrdersStore(ctx context.Context, in *DeleteOrdersStoreRequest, opts ...grpc.CallOption) (*DeleteOrdersStoreResponse, error)
	ListOrdersStore(ctx context.Context, in *ListOrdersStoresRequest, opts ...grpc.CallOption) (*ListOrdersStoresResponse, error)
}

type ordersStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersStoreServiceClient(cc grpc.ClientConnInterface) OrdersStoreServiceClient {
	return &ordersStoreServiceClient{cc}
}

func (c *ordersStoreServiceClient) CreateOrdersStore(ctx context.Context, in *CreateOrdersStoreRequest, opts ...grpc.CallOption) (*CreateOrdersStoreResponse, error) {
	out := new(CreateOrdersStoreResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.OrdersStoreService/CreateOrdersStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersStoreServiceClient) GetOrdersStore(ctx context.Context, in *GetOrdersStoreRequest, opts ...grpc.CallOption) (*GetOrdersStoreResponse, error) {
	out := new(GetOrdersStoreResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.OrdersStoreService/GetOrdersStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersStoreServiceClient) DeleteOrdersStore(ctx context.Context, in *DeleteOrdersStoreRequest, opts ...grpc.CallOption) (*DeleteOrdersStoreResponse, error) {
	out := new(DeleteOrdersStoreResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.OrdersStoreService/DeleteOrdersStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersStoreServiceClient) ListOrdersStore(ctx context.Context, in *ListOrdersStoresRequest, opts ...grpc.CallOption) (*ListOrdersStoresResponse, error) {
	out := new(ListOrdersStoresResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.OrdersStoreService/ListOrdersStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersStoreServiceServer is the server API for OrdersStoreService service.
// All implementations must embed UnimplementedOrdersStoreServiceServer
// for forward compatibility
type OrdersStoreServiceServer interface {
	CreateOrdersStore(context.Context, *CreateOrdersStoreRequest) (*CreateOrdersStoreResponse, error)
	GetOrdersStore(context.Context, *GetOrdersStoreRequest) (*GetOrdersStoreResponse, error)
	DeleteOrdersStore(context.Context, *DeleteOrdersStoreRequest) (*DeleteOrdersStoreResponse, error)
	ListOrdersStore(context.Context, *ListOrdersStoresRequest) (*ListOrdersStoresResponse, error)
	mustEmbedUnimplementedOrdersStoreServiceServer()
}

// UnimplementedOrdersStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrdersStoreServiceServer struct {
}

func (UnimplementedOrdersStoreServiceServer) CreateOrdersStore(context.Context, *CreateOrdersStoreRequest) (*CreateOrdersStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrdersStore not implemented")
}
func (UnimplementedOrdersStoreServiceServer) GetOrdersStore(context.Context, *GetOrdersStoreRequest) (*GetOrdersStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersStore not implemented")
}
func (UnimplementedOrdersStoreServiceServer) DeleteOrdersStore(context.Context, *DeleteOrdersStoreRequest) (*DeleteOrdersStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrdersStore not implemented")
}
func (UnimplementedOrdersStoreServiceServer) ListOrdersStore(context.Context, *ListOrdersStoresRequest) (*ListOrdersStoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrdersStore not implemented")
}
func (UnimplementedOrdersStoreServiceServer) mustEmbedUnimplementedOrdersStoreServiceServer() {}

// UnsafeOrdersStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersStoreServiceServer will
// result in compilation errors.
type UnsafeOrdersStoreServiceServer interface {
	mustEmbedUnimplementedOrdersStoreServiceServer()
}

func RegisterOrdersStoreServiceServer(s grpc.ServiceRegistrar, srv OrdersStoreServiceServer) {
	s.RegisterService(&OrdersStoreService_ServiceDesc, srv)
}

func _OrdersStoreService_CreateOrdersStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrdersStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersStoreServiceServer).CreateOrdersStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.OrdersStoreService/CreateOrdersStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersStoreServiceServer).CreateOrdersStore(ctx, req.(*CreateOrdersStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersStoreService_GetOrdersStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersStoreServiceServer).GetOrdersStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.OrdersStoreService/GetOrdersStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersStoreServiceServer).GetOrdersStore(ctx, req.(*GetOrdersStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersStoreService_DeleteOrdersStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrdersStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersStoreServiceServer).DeleteOrdersStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.OrdersStoreService/DeleteOrdersStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersStoreServiceServer).DeleteOrdersStore(ctx, req.(*DeleteOrdersStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersStoreService_ListOrdersStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersStoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersStoreServiceServer).ListOrdersStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.OrdersStoreService/ListOrdersStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersStoreServiceServer).ListOrdersStore(ctx, req.(*ListOrdersStoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrdersStoreService_ServiceDesc is the grpc.ServiceDesc for OrdersStoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrdersStoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "online_shop.management.v1.OrdersStoreService",
	HandlerType: (*OrdersStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrdersStore",
			Handler:    _OrdersStoreService_CreateOrdersStore_Handler,
		},
		{
			MethodName: "GetOrdersStore",
			Handler:    _OrdersStoreService_GetOrdersStore_Handler,
		},
		{
			MethodName: "DeleteOrdersStore",
			Handler:    _OrdersStoreService_DeleteOrdersStore_Handler,
		},
		{
			MethodName: "ListOrdersStore",
			Handler:    _OrdersStoreService_ListOrdersStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders_store.proto",
}
