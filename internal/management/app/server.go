package app

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc/keepalive"

	gracefulshutdown "github.com/lvlBA/online_shop/internal/graceful_shutdown"

	"github.com/lvlBA/online_shop/internal/management/app/site"
	controllersSite "github.com/lvlBA/online_shop/internal/management/controllers/site"
	"github.com/lvlBA/online_shop/internal/management/db"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

var keepAliveParams = keepalive.ServerParameters{
	MaxConnectionIdle:     time.Hour,
	MaxConnectionAge:      time.Hour,
	MaxConnectionAgeGrace: time.Minute,
	Time:                  time.Minute,
	Timeout:               time.Second * 5,
}

func Run(cfg *Config) error {
	log, err := cfg.getLogger()
	if err != nil {
		return fmt.Errorf("failed to get logger: %w", err)
	}

	grpcListener, err := cfg.getGrpcListener()
	if err != nil {
		return fmt.Errorf("failed to get listener: %w", err)
	}

	conn, err := cfg.getDatabaseConnection()
	if err != nil {
		return fmt.Errorf("failed to get db connection: %w", err)
	}

	// controllers
	dbSvc := db.New(conn)
	siteCtrl := controllersSite.New(dbSvc)
	siteApp := site.New(siteCtrl, log)

	gs := gracefulshutdown.New(&gracefulshutdown.Config{
		Ctx:           context.Background(),
		Log:           log,
		Stop:          nil,
		StopWithError: nil,
	})

	grpcSvc := cfg.getGrpcServer(gs.GrpcInterceptor)
	gs.AddStop(grpcSvc.Stop)
	api.RegisterSiteServiceServer(grpcSvc, siteApp)
	go gs.Observe()
	if err = grpcSvc.Serve(grpcListener); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}
	log.Info(gs.GetContext(), "service finished")

	return nil
}
