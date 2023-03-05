package usecase

import (
	"context"
	"inex/main/domain"
	"inex/main/internal/repository"
)

func CreateCostItem(item domain.CostItem, repo repository.InexRepo) (*domain.CostItem, error) {
	var user domain.User
	user.ID = item.UserId

	newItem, err := repo.CreateCostItem(context.Background(), item, user)

	if err != nil {
		return nil, err
	}

	return newItem, nil
}

func ReadCostItems(user domain.User, repo repository.InexRepo) (*[]domain.CostItem, error) {
	items, err := repo.ReadCostItems(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return &items, nil
}

func UpdateCostItem(item domain.CostItem, repo repository.InexRepo) (*domain.CostItem, error) {
	newItem, err := repo.UpdateCostItem(context.Background(), item)
	if err != nil {
		return nil, err
	}

	return newItem, nil
}

func DeleteCostItem(item domain.CostItem, repo repository.InexRepo) error {
	err := repo.DeleteCostItem(context.Background(), item)
	if err != nil {
		return err
	}

	return nil
}
