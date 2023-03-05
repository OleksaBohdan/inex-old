package usecase

import (
	"context"
	"inex/main/domain"
	"inex/main/internal/repository"
	"math"
)

func CreateIncome(income domain.Income, repo repository.InexRepo) (*domain.Income, error) {

	income.Value = math.Trunc(income.Value*100) / 100

	newIncome, err := repo.CreateIncome(context.Background(), income)
	if err != nil {
		return nil, err
	}

	return newIncome, nil
}

//func UpdateIncome() (*domain.Income, error) {
//
//}
