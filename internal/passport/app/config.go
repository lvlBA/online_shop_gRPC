package app

type Config struct {
	GrpcAddr string `json:"grpc_addr" yaml:"grpc_addr" env:"GRPC_ADDR" envDefault:":9090"`
	DbHost   string `json:"db_host"   yaml:"db_host"   env:"DB_HOST"   envDefault:"postgres://db:db@localhost:5478/db"`
}
