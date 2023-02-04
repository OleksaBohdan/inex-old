package app

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"inex/main/configs"
	"log"
)

func init() {
	fmt.Println("run migrate")

	pg, err := configs.NewPG()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", pg.URL)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		pg.Migrations, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		log.Println("migrate:", err)
	}
}
