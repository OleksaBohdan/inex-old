package usecase

import (
	"context"
	"inex/main/domain"
	"inex/main/internal/repository"
)

func CreateIncomeItem(item domain.IncomeItem, repo repository.InexRepo) (*domain.IncomeItem, error) {
	var user domain.User
	user.ID = item.UserId

	newItem, err := repo.CreateIncomeItem(context.Background(), item, user)

	if err != nil {
		return nil, err
	}

	return newItem, nil
}

func ReadIncomeItems(user domain.User, repo repository.InexRepo) (*[]domain.IncomeItem, error) {
	items, err := repo.ReadIncomeItems(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return &items, nil
}

func UpdateIncomeItem(item domain.IncomeItem, repo repository.InexRepo) (*domain.IncomeItem, error) {
	newItem, err := repo.UpdateIncomeItem(context.Background(), item)
	if err != nil {
		return nil, err
	}

	return newItem, nil
}

func DeleteIncomeItem(item domain.IncomeItem, repo repository.InexRepo) error {
	err := repo.DeleteIncomeItem(context.Background(), item)
	if err != nil {
		return err
	}

	return nil
}
