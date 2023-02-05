package repository

import (
	"context"
	"fmt"
	"inex/main/database/sqlQueries"
	"inex/main/domain"
	"inex/main/pkg/postgres"
)

type (
	InexRepo struct {
		*postgres.Postgres
	}
)

func New(pg *postgres.Postgres) *InexRepo {
	return &InexRepo{pg}
}

func (r InexRepo) CreateUser(ctx context.Context, user domain.User) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.CreateUser, user.Name, user.Surname, user.Email, user.PasswordHash)
	if err != nil {
		return fmt.Errorf("InexRepo - CreateUser - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	return nil
}

func (r InexRepo) ReadUser(ctx context.Context, user domain.User) (*domain.User, error) {
	rows, err := r.Pool.Query(ctx, sqlQueries.ReadUser, user.ID)
	if err != nil {
		return nil, fmt.Errorf("InexRepo - ReadUser - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var u domain.User

	return &u, nil
}

func (r InexRepo) UpdateUser(ctx context.Context, user domain.User) error {

	return nil
}

func (r InexRepo) DeleteUser(ctx context.Context, user domain.User) error {

	return nil
}
