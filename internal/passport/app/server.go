package app

import (
	"context"
	"fmt"
	grpcinterceptors "github.com/lvlBA/online_shop/internal/grpc_interceptors"
	garbagecollector "github.com/lvlBA/online_shop/internal/passport/garbage_collector"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc/keepalive"

	gracefulshutdown "github.com/lvlBA/online_shop/internal/graceful_shutdown"
	appAuth "github.com/lvlBA/online_shop/internal/passport/app/auth"
	appResource "github.com/lvlBA/online_shop/internal/passport/app/resource"
	appUser "github.com/lvlBA/online_shop/internal/passport/app/user"
	controllersAuth "github.com/lvlBA/online_shop/internal/passport/controllers/auth"
	controllersResource "github.com/lvlBA/online_shop/internal/passport/controllers/resource"
	controllersUser "github.com/lvlBA/online_shop/internal/passport/controllers/user"
	"github.com/lvlBA/online_shop/internal/passport/db"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
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
	getUserMetaIntr := grpcinterceptors.NewGetUserMeta()
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
	userCtrl := controllersUser.New(dbSvc)
	userApp := appUser.New(userCtrl, log)

	resourceCtrl := controllersResource.New(dbSvc)
	resourceApp := appResource.New(resourceCtrl, log)

	authCtrl := controllersAuth.New(dbSvc)
	authApp := appAuth.New(&appAuth.Config{
		Log:          log,
		CtrlAuth:     authCtrl,
		CtrlResource: resourceCtrl,
		CtrlUser:     userCtrl,
	})

	// garbage collectors
	gcToken := garbagecollector.NewToken(&garbagecollector.Config{
		Log:     log,
		Expired: cfg.TokenExpired,
		DB:      dbSvc,
		Timeout: time.Hour * 24,
	})

	// GRPC register
	grpcSvc := cfg.getGrpcServer(gs.GrpcInterceptor, getUserMetaIntr.GrpcInterceptor)
	gs.AddStop(grpcSvc.Stop)
	api.RegisterUserServiceServer(grpcSvc, userApp)
	api.RegisterResourceServiceServer(grpcSvc, resourceApp)
	api.RegisterAuthServiceServer(grpcSvc, authApp)

	gs.GetWG().Add(1)
	go gcToken.Observe(gs.GetContext(), gs.GetWG())
	go gs.Observe()

	if err = grpcSvc.Serve(grpcListener); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	log.Info(gs.GetContext(), "service finished")

	return nil
}
