package configs

import (
	"os"
)

type Config struct {
}

type PG struct {
	URL string `env-required:"true" env:"PG_URL"`
}

func NewPG() (*PG, error) {
	var pg PG

	pg.URL = os.Getenv("PG_URL")

	if len(pg.URL) == 0 {
		pg.URL = "postgres://postgres:inex1234@localhost:5432/postgres"
	}

	return &pg, nil
}
