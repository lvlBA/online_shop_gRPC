package main

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v7"

	"github.com/lvlBA/online_shop/internal/passport/app"
)

func main() {
	cfg := &app.Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("failed to parse config: %s", err)
		os.Exit(1)
	}

	if err := app.Run(cfg); err != nil {
		fmt.Printf("failed to run: %s", err)
		os.Exit(2)
	}
}
