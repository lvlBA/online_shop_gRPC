package app

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc/keepalive"

	gracefulshutdown "github.com/lvlBA/online_shop/internal/graceful_shutdown"
	grpcinterceptors "github.com/lvlBA/online_shop/internal/grpc_interceptors"
	appGoods "github.com/lvlBA/online_shop/internal/storage/app/goods"
	controllersGoods "github.com/lvlBA/online_shop/internal/storage/controllers/goods"
	"github.com/lvlBA/online_shop/internal/storage/db"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
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

	// interceptors
	getUserMetaInter := grpcinterceptors.NewGetUserMeta()
	gs := gracefulshutdown.New(&gracefulshutdown.Config{
		Ctx:  context.Background(),
		Log:  log,
		Stop: nil,
	})

	grpcListener, err := cfg.getGrpcListener()
	if err != nil {
		return fmt.Errorf("failed to get listener: %w", err)
	}
	defer func() {
		if err := grpcListener.Close(); err != nil {
			log.Error(gs.GetContext(), " failed to close listener %w", err)
		}
	}()

	conn, err := cfg.getDatabaseConnection()
	if err != nil {
		return fmt.Errorf("failed to get db connection: %w", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Error(gs.GetContext(), " failed to close db connection %w", err)
		}
	}()

	// controllers
	dbSvc := db.New(conn)
	goodsCtrl := controllersGoods.New(dbSvc)
	goodsApp := appGoods.New(goodsCtrl, log)

	// GRPC register
	grpcSvc := cfg.getGrpcServer(gs.GrpcInterceptor, getUserMetaInter.GrpcInterceptor)
	gs.AddStop(grpcSvc.Stop)
	api.RegisterGoodsServiceServer(grpcSvc, goodsApp)

	if err = grpcSvc.Serve(grpcListener); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	log.Info(gs.GetContext(), "service finished")

	return nil
}