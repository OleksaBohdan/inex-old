package app

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
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

	// Create
	//var user domain.User
	//inex := repository.New(pg)
	//user.Name = "Julia"
	//user.Surname = "Meln"
	//user.Email = "test0@gmail.com"
	//user.PasswordHash = "kgfsdgfcsdgc27f38tf723gfuewifgskudgskjfsdkfjdsdfdsjada"
	//
	//err = inex.CreateUser(context.Background(), user)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Rux - inex.CreateUser: %w", err))
	//}

	// Update
	//var user domain.User
	//inex := repository.New(pg)
	//id, err := uuid.FromString("cd249fd9-80f9-40b1-8d53-4e7585fc99b7")
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Rux - uuid: %w", err))
	//}
	//user.ID = id
	//user.Name = "Bohdan_Updated"
	//user.Surname = "Oleksa"
	//user.Email = "test3@gmail.com"
	//user.PasswordHash = "kgfsdgfcsdgc27f38tf723gfuewifgskudgskjfsdkfjdsbfdsjada"
	//
	//err = inex.UpdateUser(context.Background(), user)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Rux - inex.Update: %w", err))
	//}
	//
	//fmt.Println("Done")

	// Delete
	var user domain.User
	inex := repository.New(pg)
	id, err := uuid.FromString("a42a7a77-d06e-410f-9630-2a39913cb34d")
	if err != nil {
		l.Fatal(fmt.Errorf("app - Rux - uuid: %w", err))
	}
	user.ID = id

	//err = inex.DeleteUser(context.Background(), user)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Rux - inex.Delete: %w", err))
	//}

	//NOTES

	//var note domain.Note
	//note.Text = "Lorem lorem lorem text updated"

	//err = inex.CreateNote(context.Background(), note, user)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Rux - inex.Delete: %w", err))
	//}

	//err = inex.UpdateNote(context.Background(), note, user)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Rux - inex.UpdateNote: %w", err))
	//}

	// read Note
	//n, err := inex.ReadNote(context.Background(), user)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Rux - inex.ReadNote: %w", err))
	//}
	//fmt.Println("note:", n.Text)

	err = inex.DeleteNote(context.Background(), user)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Rux - inex.DeleteNote: %w", err))
	}

}

