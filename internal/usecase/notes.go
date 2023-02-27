package usecase

import (
	"fmt"
	"inex/main/domain"
	"inex/main/internal/repository"
)

func CreateNote(note domain.Note, repo repository.InexRepo) error {
	fmt.Println("usecase - CreateNote")

	return nil
}
