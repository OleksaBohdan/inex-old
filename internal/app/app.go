package app

import (
	"context"
	"fmt"
	"inex/main/config"
	"inex/main/domain"
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

	// test seeding database

	var user domain.User
	inex := repository.New(pg)
	user.Name = "Bohdan"
	user.Surname = "Oleksa"
	user.Email = "test@gmail.com"
	user.PasswordHash = "kgfsdgfcsdgc27f38tf723gfuewifgskudgskjfsdkfjdsbfdsjf"

	err = inex.CreateUser(context.Background(), user)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Rux - inex.CreateUser: %w", err))
	}
}
