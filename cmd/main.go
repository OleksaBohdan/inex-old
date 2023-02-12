package main

import (
	"inex/main/config"
	"inex/main/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	app.Run(cfg)
}
