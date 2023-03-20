package grpcinterceptors

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"reflect"
)

type SetUserMeta struct {
}

func NewSetUserMeta() *SetUserMeta {
	return &SetUserMeta{}
}

func (c *SetUserMeta) GrpcInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := make(metadata.MD)
	if valueI := ctx.Value(xRequestIdKey); valueI != nil {
		value, ok := valueI.(string)
		if !ok {
			return fmt.Errorf("failed to type assertion, want 'string' != got '%s'\n", reflect.TypeOf(valueI))
		}
		md.Set(xRequestIdKey, value)
	}

	if valueI := ctx.Value(tokenKey); valueI != nil {
		value, ok := valueI.([]byte)
		if !ok {
			return fmt.Errorf("failed to type assertion, want '[]byte' != got '%s'\n", reflect.TypeOf(valueI))
		}
		md.Set(tokenKey, string(value))
	}

	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}
