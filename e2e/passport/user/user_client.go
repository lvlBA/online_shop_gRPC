package user

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

type SiteClientImpl struct {
	api.UserServiceClient
}

type Config struct {
	Addr string `env:"ADDR" envDefault:"localhost:9092"`
}

func New(ctx context.Context, cfg *Config) (*SiteClientImpl, error) {
	conn, err := grpc.DialContext(ctx, cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error create connection: %s", err)
	}
	return &SiteClientImpl{
		UserServiceClient: api.NewUserServiceClient(conn),
	}, nil
}
