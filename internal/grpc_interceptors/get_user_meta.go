package grpcinterceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GetUserMeta struct{}

func NewGetUserMeta() *GetUserMeta {
	return &GetUserMeta{}
}

func (c *GetUserMeta) GrpcInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return handler(ctx, req)
	}
	if values, exists := md[xRequestIdKey]; exists {
		if len(values) > 0 {
			ctx = context.WithValue(ctx, xRequestIdKey, values[0])
		}
	}

	if values, exists := md[tokenKey]; exists {
		if len(values) > 0 {
			ctx = context.WithValue(ctx, tokenKey, values[0])
		}
	}

	return handler(ctx, req)
}
