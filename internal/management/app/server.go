package app

import (
	"fmt"
	"github.com/lvlBA/online_shop/internal/management/app/site"
	controllersSite "github.com/lvlBA/online_shop/internal/management/controllers/site"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
	"net"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/lvlBA/online_shop/internal/management/db"
)

var keepAliveParams = keepalive.ServerParameters{
	MaxConnectionIdle:     time.Hour,
	MaxConnectionAge:      time.Hour,
	MaxConnectionAgeGrace: time.Minute,
	Time:                  time.Minute,
	Timeout:               time.Second * 5,
}

func Run(cfg *Config) error {
	grpcSvc := grpc.NewServer(
		grpc.KeepaliveParams(keepAliveParams),
	)

	grpcListener, err := net.Listen("tcp", cfg.GrpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen grpc address: %w", err)
	}

	pgConn, err := sqlx.Connect("pgx", cfg.DbHost)
	if err != nil {
		return fmt.Errorf("failed to connecto to db: %w", err)
	}

	dbSvc := db.New(pgConn)
	siteCtrl := controllersSite.New(dbSvc)
	siteApp := site.New(siteCtrl)

	api.RegisterSiteServiceServer(grpcSvc, siteApp)

	if err = grpcSvc.Serve(grpcListener); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}

	return nil
}
