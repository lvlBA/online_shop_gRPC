package grpcinterceptors

import (
	"context"
	"google.golang.org/grpc"
)

type GetUserMeta struct{}

func NewGetUserMeta() *GetUserMeta {
	return &GetUserMeta{}
}

func (c *GetUserMeta) GrpcInterceptor(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx = metaToContext(ctx, XRequestIdKey, TokenKey)

	return handler(ctx, req)
}
