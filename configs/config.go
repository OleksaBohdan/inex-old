package configs

import (
	"os"
)

type Config struct {
}

type PG struct {
	URL        string `env-required:"true" env:"PG_URL"`
	Migrations string `env-required:"true" env:"PG_MIGRATIONS"`
}

func NewPG() (*PG, error) {
	var pg PG

	pg.URL = os.Getenv("PG_URL")
	pg.Migrations = os.Getenv("PG_MIGRATIONS")

	if len(pg.URL) == 0 {
		pg.URL = "postgres://postgres:inex1234@localhost:5432/postgres?sslmode=disable"
	}

	if len(pg.Migrations) == 0 {
		pg.Migrations = "file://../database/migrations"
	}

	return &pg, nil
}
