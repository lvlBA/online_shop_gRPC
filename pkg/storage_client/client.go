package storage_client

import (
	"context"
	"fmt"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Addr string `env:"ADDR" envDefault:"localhost:9096"`
}

type ClientImpl struct {
	api.CargoServiceClient
	api.GoodsServiceClient
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
		CargoServiceClient: api.NewCargoServiceClient(conn),
		GoodsServiceClient: api.NewGoodsServiceClient(conn),
	}, nil
}
