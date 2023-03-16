package app

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc/keepalive"

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

	grpcSvc := cfg.getGrpcServer()
	api.RegisterUserServiceServer(grpcSvc, userApp)
	api.RegisterResourceServiceServer(grpcSvc, resourceApp)
	api.RegisterAuthServiceServer(grpcSvc, authApp)
	if err = grpcSvc.Serve(grpcListener); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}

	return nil
}
