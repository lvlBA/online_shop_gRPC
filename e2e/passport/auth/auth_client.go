package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

type SiteClientImpl struct {
	api.AuthServiceClient
	api.ResourceServiceClient
	api.UserServiceClient
}

type Config struct {
	Addr string `env:"ADDR" envDefault:"localhost:9091"`
}

func New(ctx context.Context, cfg *Config) (*SiteClientImpl, error) {
	conn, err := grpc.DialContext(ctx, cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error create connection: %s", err)
	}
	return &SiteClientImpl{
		AuthServiceClient:     api.NewAuthServiceClient(conn),
		ResourceServiceClient: api.NewResourceServiceClient(conn),
		UserServiceClient:     api.NewUserServiceClient(conn),
	}, nil
}
