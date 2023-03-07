package app

import (
	"fmt"
	"net"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/lvlBA/online_shop/internal/management/app/site"
	controllersSite "github.com/lvlBA/online_shop/internal/management/controllers/site"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/pkg/logger"
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
	zapLoggerCfg := zap.NewProductionConfig()
	zapLoggerCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	zapLoggerCfg.DisableCaller = true
	zapLoggerCfg.DisableStacktrace = true
	zapLogger, err := zapLoggerCfg.Build()
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}
	log := logger.NewZapLogger(zapLogger.Sugar())

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
	siteApp := site.New(siteCtrl, log)

	api.RegisterSiteServiceServer(grpcSvc, siteApp)

	if err = grpcSvc.Serve(grpcListener); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}

	return nil
}
