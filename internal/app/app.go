package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"inex/main/config"
	"inex/main/internal/controller/http/routes"
	"inex/main/internal/repository"
	"inex/main/pkg/logger"
	"inex/main/pkg/postgres"
)

func Run(cfg *config.Config) {
	fmt.Println("Run app...")
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	//seeder.SeedDatabase(cfg)

	inexRepo := repository.New(pg)

	e := echo.New()
	routes.RegisterRoutes(e, inexRepo)

	e.Logger.Fatal(e.Start(cfg.HTTP.Port))
}
