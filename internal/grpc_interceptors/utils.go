package grpcinterceptors

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func metaToContext(ctx context.Context, keys ...string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for _, key := range keys {
			if values, exists := md[key]; exists {
				if len(values) > 0 {
					ctx = context.WithValue(ctx, key, values[0])
				}
			}
		}
	}

	return ctx
}

func contextToMeta(ctx context.Context, keys ...string) context.Context {
	md := make(metadata.MD)
	for _, key := range keys {
		if valueI := ctx.Value(key); valueI != nil {
			switch value := valueI.(type) {
			case string:
				md.Set(key, value)
			case []byte:
				md.Set(key, string(value))
			default:
				continue
			}
		}
	}

	return metadata.NewOutgoingContext(ctx, md)
}
