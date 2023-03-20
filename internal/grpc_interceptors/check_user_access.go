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

func (c *CheckUserAccess) GrpcInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if _, err := c.cli.CheckUserAccess(ctx, &api.CheckUserAccessRequest{
		ResourceId: info.FullMethod,
	}); err != nil {
		if status.Code(err) == codes.Unauthenticated {
			return nil, err
		}
		c.log.Error(ctx, "failed to check user access", err, "req", req)

		return nil, status.Error(codes.Internal, "error check user access")
	}

	return handler(ctx, req)
}
