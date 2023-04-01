package app

import (
	"net"
	"time"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"

	"github.com/lvlBA/online_shop/internal/bootstrap"
	"github.com/lvlBA/online_shop/pkg/logger"
)

type Config struct {
	GrpcAddr         string        `json:"grpc_addr"     yaml:"grpc_addr"     env:"GRPC_ADDR"     envDefault:":9096"`
	GrpcPassportAddr string        `json:"passport_grpc_addr" yaml:"passport_grpc_addr" env:"PASSPORT_GRPC_ADDR" envDefault:":9092"`
	DbHost           string        `json:"db_host"       yaml:"db_host"       env:"DB_HOST"       envDefault:"postgres://db:db@localhost:5478/db"`
	LogLevel         string        `json:"log_level"     yaml:"log_level"     env:"LOG_LEVEL"     envDefault:"error"`
	TokenExpired     time.Duration `json:"token_expired" yaml:"token_expired" env:"DURATION"      envDefault:"10m"`
}

func (c *Config) getLogger() (logger.Logger, error) {
	return bootstrap.InitLogger(&bootstrap.ConfigLogger{
		Level:             c.LogLevel,
		DisableCaller:     true,
		DisableStacktrace: true,
	})
}

func (c *Config) getDatabaseConnection() (*sqlx.DB, error) {
	return sqlx.Connect("pgx", c.DbHost)
}

func (c *Config) getGrpcListener() (net.Listener, error) {
	return net.Listen("tcp", c.GrpcAddr)
}

func (c *Config) getGrpcServer(inters ...grpc.UnaryServerInterceptor) *grpc.Server {
	return grpc.NewServer(
		grpc.KeepaliveParams(keepAliveParams),
		grpc.ChainUnaryInterceptor(inters...),
	)
}
