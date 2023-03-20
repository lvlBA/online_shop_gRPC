package auth

type Config struct {
	Addr string `env:"ADDR" envDefault:"localhost:9092"`
}
