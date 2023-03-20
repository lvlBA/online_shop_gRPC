package passportclient

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

type Config struct {
	Addr string `env:"ADDR" envDefault:"localhost:9091"`
}

type ClientImpl struct {
	api.AuthServiceClient
	api.ResourceServiceClient
	api.UserServiceClient
}

func New(ctx context.Context, cfg *Config, intrs ...grpc.UnaryClientInterceptor) (*ClientImpl, error) {
	conn, err := grpc.DialContext(
		ctx,
		cfg.Addr,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithChainUnaryInterceptor(intrs...),
	)
	if err != nil {
		return nil, fmt.Errorf("error create connection: %s", err)
	}
	return &ClientImpl{
		AuthServiceClient:     api.NewAuthServiceClient(conn),
		ResourceServiceClient: api.NewResourceServiceClient(conn),
		UserServiceClient:     api.NewUserServiceClient(conn),
	}, nil
}
