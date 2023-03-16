package resource

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

type SiteClientImpl struct {
	api.ResourceServiceClient
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
		ResourceServiceClient: api.NewResourceServiceClient(conn),
	}, nil
}
