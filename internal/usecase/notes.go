package usecase

import (
	"context"
	"inex/main/domain"
	"inex/main/internal/repository"
)

func CreateNote(note domain.Note, repo repository.InexRepo) (*domain.Note, error) {
	var user domain.User
	user.ID = note.UserId

	newNote, err := repo.ReadNote(context.Background(), user)
	if newNote != nil {
		newNote, err = repo.UpdateNote(context.Background(), note, user)
		if err != nil {
			return nil, err
		}
		return newNote, nil
	}

	newNote, err = repo.CreateNote(context.Background(), note)
	if err != nil {
		return nil, err
	}

	return newNote, nil
}
