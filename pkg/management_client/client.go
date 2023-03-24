package managementclient

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

type Config struct {
	Addr string `env:"ADDR" envDefault:"localhost:9090"`
}

type ClientImpl struct {
	api.SiteServiceClient
	api.RegionServiceClient
	api.LocationServiceClient
	api.WarehouseServiceClient
	api.OrdersStoreServiceClient
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
		SiteServiceClient:        api.NewSiteServiceClient(conn),
		RegionServiceClient:      api.NewRegionServiceClient(conn),
		LocationServiceClient:    api.NewLocationServiceClient(conn),
		WarehouseServiceClient:   api.NewWarehouseServiceClient(conn),
		OrdersStoreServiceClient: api.NewOrdersStoreServiceClient(conn),
	}, nil
}
