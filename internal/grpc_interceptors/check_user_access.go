package grpcinterceptors

import (
	"context"
	"github.com/lvlBA/online_shop/pkg/logger"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckUserAccess struct {
	log logger.Logger
	cli api.AuthServiceClient
}

func New(log logger.Logger, cli api.AuthServiceClient) *CheckUserAccess {
	return &CheckUserAccess{
		log: log,
		cli: cli,
	}
}

func (c *CheckUserAccess) GrpcInterceptor(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	ctxWithMeta := metaToContext(ctx, XRequestIdKey, TokenKey)
	_, err := c.cli.CheckUserAccess(contextToMeta(ctxWithMeta, XRequestIdKey, TokenKey), &api.CheckUserAccessRequest{Resource: info.FullMethod})
	if err != nil {
		if status.Code(err) == codes.Unauthenticated {
			return nil, err
		}
		c.log.Error(ctxWithMeta, "failed to check user access", err, "req", req)

		return nil, status.Error(codes.Internal, "error check user access")
	}

	return handler(ctxWithMeta, req)
}
