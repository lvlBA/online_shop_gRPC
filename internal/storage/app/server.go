package app

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc/keepalive"

	gracefulshutdown "github.com/lvlBA/online_shop/internal/graceful_shutdown"
	grpcInterceptors "github.com/lvlBA/online_shop/internal/grpc_interceptors"
	appCargo "github.com/lvlBA/online_shop/internal/storage/app/cargo"
	appGoods "github.com/lvlBA/online_shop/internal/storage/app/goods"
	controllersCargo "github.com/lvlBA/online_shop/internal/storage/controllers/cargo"
	controllersGoods "github.com/lvlBA/online_shop/internal/storage/controllers/goods"
	"github.com/lvlBA/online_shop/internal/storage/db"
	passportClient "github.com/lvlBA/online_shop/pkg/passport_client"
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

	cargoCtrl := controllersCargo.New(dbSvc)
	cargoApp := appCargo.New(cargoCtrl, log)

	passportCli, err := passportClient.New(gs.GetContext(), &passportClient.Config{
		Addr: cfg.GrpcPassportAddr,
	})

	if err != nil {
		return fmt.Errorf("failed to create passport client: %w", err)
	}

	userAuthInter := grpcInterceptors.New(log, passportCli)
	grpcSvc := cfg.getGrpcServer(gs.GrpcInterceptor, userAuthInter.GrpcInterceptor)

	gs.AddStop(grpcSvc.Stop)
	// GRPC register

	api.RegisterGoodsServiceServer(grpcSvc, goodsApp)
	api.RegisterCargoServiceServer(grpcSvc, cargoApp)

	if err = grpcSvc.Serve(grpcListener); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	log.Info(gs.GetContext(), "service finished")

	return nil
}
