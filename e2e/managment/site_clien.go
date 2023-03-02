package managment

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"

	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

type SiteClientImpl struct {
	api.SiteServiceClient
}

type Config struct {
	Addr string "env:\"ADDR\" envDefault: \"Localhost: 9090\""
}

func New(ctx context.Context, cfg *Config) (*SiteClientImpl, error) {
	conn, err := grpc.DialContext(ctx, cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error create connection: %s", err)
	}
	return &SiteClientImpl{
		SiteServiceClient: api.NewSiteServiceClient(conn),
	}, nil
}
