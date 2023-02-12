package app

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"inex/main/config"
	"log"
)

func init() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.PG.URL)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		cfg.PG.Migrations, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		log.Println("migrate:", err)
	}
}
