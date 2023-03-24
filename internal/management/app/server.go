package app

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc/keepalive"

	gracefulShutdown "github.com/lvlBA/online_shop/internal/graceful_shutdown"
	grpcInterceptors "github.com/lvlBA/online_shop/internal/grpc_interceptors"
	"github.com/lvlBA/online_shop/internal/management/app/location"
	"github.com/lvlBA/online_shop/internal/management/app/orders_store"
	"github.com/lvlBA/online_shop/internal/management/app/region"
	"github.com/lvlBA/online_shop/internal/management/app/site"
	"github.com/lvlBA/online_shop/internal/management/app/warehouse"
	controllersLocation "github.com/lvlBA/online_shop/internal/management/controllers/location"
	controllersOrdersStore "github.com/lvlBA/online_shop/internal/management/controllers/orders_store"
	controllersRegion "github.com/lvlBA/online_shop/internal/management/controllers/region"
	controllersSite "github.com/lvlBA/online_shop/internal/management/controllers/site"
	controllersWarehouse "github.com/lvlBA/online_shop/internal/management/controllers/warehouse"
	"github.com/lvlBA/online_shop/internal/management/db"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
	passportClient "github.com/lvlBA/online_shop/pkg/passport_client"
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

	gs := gracefulShutdown.New(&gracefulShutdown.Config{
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

	siteCtrl := controllersSite.New(dbSvc)
	siteApp := site.New(siteCtrl, log)

	regionCtrl := controllersRegion.New(dbSvc)
	regionApp := region.New(regionCtrl, log)

	locationCtrl := controllersLocation.New(dbSvc)
	locationApp := location.New(locationCtrl, log)

	warehouseCtrl := controllersWarehouse.New(dbSvc)
	warehouseApp := warehouse.New(warehouseCtrl, log)

	ordersStoreCtrl := controllersOrdersStore.New(dbSvc)
	ordersStoreApp := orders_store.New(ordersStoreCtrl, log)

	passportCli, err := passportClient.New(gs.GetContext(), &passportClient.Config{
		Addr: cfg.GrpcPassportAddr,
	})
	if err != nil {
		return fmt.Errorf("failed to create passport client: %w", err)
	}

	userAuthInter := grpcInterceptors.New(log, passportCli)
	grpcSvc := cfg.getGrpcServer(gs.GrpcInterceptor, userAuthInter.GrpcInterceptor)

	gs.AddStop(grpcSvc.Stop)
	api.RegisterSiteServiceServer(grpcSvc, siteApp)
	api.RegisterRegionServiceServer(grpcSvc, regionApp)
	api.RegisterLocationServiceServer(grpcSvc, locationApp)
	api.RegisterWarehouseServiceServer(grpcSvc, warehouseApp)
	api.RegisterOrdersStoreServiceServer(grpcSvc, ordersStoreApp)
	go gs.Observe()
	if err = grpcSvc.Serve(grpcListener); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}
	log.Info(gs.GetContext(), "service finished")

	return nil
}
