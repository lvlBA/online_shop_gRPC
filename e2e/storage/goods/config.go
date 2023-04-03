package storage

type Config struct {
	Addr         string `env:"ADDR" envDefault:"localhost:9096"`
	PassportAddr string `env:"PASSPORT_ADDR" envDefault:"localhost:9092"`
}
