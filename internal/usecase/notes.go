package usecase

import (
	"context"
	"fmt"
	"inex/main/domain"
	"inex/main/internal/repository"
)

func CreateNote(note domain.Note, repo repository.InexRepo) (*domain.Note, error) {
	fmt.Println("usecase - CreateNote")

	// TODO if note exists – update note either create a new one
	var user domain.User
	user.ID = note.UserId
	not, err := repo.ReadNote(context.Background(), user)
	if err != nil {
		return nil, err
	}

	fmt.Println("not: ", not)

	//n, err := repo.CreateNote(context.Background(), note)
	//if err != nil {
	//	return nil, err
	//}

	return &note, nil
}
