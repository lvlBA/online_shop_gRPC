package grpcinterceptors

import (
	"context"

	"google.golang.org/grpc"
)

type SetUserMeta struct{}

func NewSetUserMeta() *SetUserMeta {
	return &SetUserMeta{}
}

func (c *SetUserMeta) GrpcInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = contextToMeta(ctx, XRequestIdKey, TokenKey)

	return invoker(ctx, method, req, reply, cc, opts...)
}
