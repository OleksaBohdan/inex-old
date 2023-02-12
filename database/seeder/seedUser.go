package seeder

import (
	"context"
	"fmt"
	"github.com/jaswdr/faker"
	"inex/main/domain"
	"inex/main/internal/repository"
	"inex/main/pkg/postgres"
)

func seedUsers(pg *postgres.Postgres) error {
	inex := repository.New(pg)
	fake := faker.New()

	var user domain.User

	for i := 0; i < 5; i++ {
		user.Name = fake.Person().Name()
		user.Surname = fake.Person().LastName()
		user.Email = fake.Internet().Email()
		user.PasswordHash = fake.Hash().SHA256()

		err := inex.CreateUser(context.Background(), user)
		if err != nil {
			return fmt.Errorf("seeder - seedUser - CreateUser %w", err)
		}

	}

	return nil
}

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
//var user domain.User
//inex := repository.New(pg)
//id, err := uuid.FromString("a42a7a77-d06e-410f-9630-2a39913cb34d")
//if err != nil {
//l.Fatal(fmt.Errorf("app - Rux - uuid: %w", err))
//}
//user.ID = id

//err = inex.DeleteUser(context.Background(), user)
//if err != nil {
//	l.Fatal(fmt.Errorf("app - Rux - inex.Delete: %w", err))
//}
