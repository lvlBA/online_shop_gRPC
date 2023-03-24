package location

type Config struct {
	Addr         string `env:"ADDR" envDefault:"localhost:9090"`
	PassportAddr string `env:"PASSPORT_ADDR" envDefault:"localhost:9092"`
}
