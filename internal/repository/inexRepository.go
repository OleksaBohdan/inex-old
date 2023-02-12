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

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) ReadUser(ctx context.Context, user domain.User) (*domain.User, error) {
	rows, err := r.Pool.Query(ctx, sqlQueries.ReadUser, user.ID)
	if err != nil {
		return nil, fmt.Errorf("InexRepo - ReadUser - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var u domain.User
	err = rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Email, &u.PasswordHash, &u.RegisteredAt, &u.LastVisit)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r InexRepo) UpdateUser(ctx context.Context, user domain.User) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.UpdateUser, user.Name, user.Surname, user.Email, user.PasswordHash, user.ID)
	if err != nil {
		return fmt.Errorf("InexRepo - UpdateUser - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) DeleteUser(ctx context.Context, user domain.User) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.DeleteUser, user.ID)
	if err != nil {
		return fmt.Errorf("InexRepo - DeleteUser - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) CreateNote(ctx context.Context, note domain.Note, user domain.User) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.CreateNote, user.ID, note.Text)
	if err != nil {
		return fmt.Errorf("InexRepo - CreateNote - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) ReadNote(ctx context.Context, user domain.User) (*domain.Note, error) {
	rows, err := r.Pool.Query(ctx, sqlQueries.ReadNote, user.ID)
	if err != nil {
		return nil, fmt.Errorf("InexRepo - ReadNote - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var note domain.Note
	err = rows.Scan(&note.ID, &note.UserId, &note.Text)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r InexRepo) UpdateNote(ctx context.Context, note domain.Note, user domain.User) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.UpdateNote, note.Text, user.ID)
	if err != nil {
		return fmt.Errorf("InexRepo - UpdateNote - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) DeleteNote(ctx context.Context, user domain.User) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.DeleteNote, user.ID)
	if err != nil {
		return fmt.Errorf("InexRepo - DeleteNote - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) CreateIncomeItem(ctx context.Context, item domain.IncomeItem, user domain.User) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.CreateIncomeItem, item.Name, user.ID)
	if err != nil {
		return fmt.Errorf("InexRepo - CreateIncomeItem - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) ReadIncomeItems(ctx context.Context, user domain.User) ([]domain.IncomeItem, error) {
	rows, err := r.Pool.Query(ctx, sqlQueries.ReadIncomeItems, user.ID)
	if err != nil {
		return nil, fmt.Errorf("InexRepo - ReadIncomeItems - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var items []domain.IncomeItem

	for rows.Next() {
		if rows.Err() != nil {
			return nil, rows.Err()
		}

		var item domain.IncomeItem
		err = rows.Scan(&item.ID, &item.UserId, &item.Name, &item.Rang)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (r InexRepo) UpdateIncomeItem(ctx context.Context, item domain.IncomeItem) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.UpdateIncomeItem, item.Name, item.ID)
	if err != nil {
		return fmt.Errorf("InexRepo - UpdateIncomeItem - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) DeleteIncomeItem(ctx context.Context, item domain.IncomeItem) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.DeleteIncomeItem, item.ID)
	if err != nil {
		return fmt.Errorf("InexRepo - DeleteIncomeItem - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) CreateCostItem(ctx context.Context, item domain.CostItem, user domain.User) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.CreateCostItem, item.Name, user.ID)
	if err != nil {
		return fmt.Errorf("InexRepo - CreateCostItem - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) ReadCostItems(ctx context.Context, user domain.User) ([]domain.CostItem, error) {
	rows, err := r.Pool.Query(ctx, sqlQueries.ReadCostItems, user.ID)
	if err != nil {
		return nil, fmt.Errorf("InexRepo - ReadCostItems - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var items []domain.CostItem

	for rows.Next() {
		if rows.Err() != nil {
			return nil, rows.Err()
		}

		var item domain.CostItem
		err = rows.Scan(&item.ID, &item.UserId, &item.Name, &item.Rang)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (r InexRepo) UpdateCostItem(ctx context.Context, item domain.CostItem) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.UpdateCostItem, item.Name, item.ID)
	if err != nil {
		return fmt.Errorf("InexRepo - UpdateCostItem - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) DeleteCostItem(ctx context.Context, item domain.CostItem) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.DeleteCostItem, item.ID)
	if err != nil {
		return fmt.Errorf("InexRepo - DeleteCostItem - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) CreateIncome(ctx context.Context, income domain.Income, incomeItem domain.IncomeItem, user domain.User) error {
	rows, err := r.Pool.Query(ctx, sqlQueries.CreateIncome, user.ID, incomeItem.ID, income.Value)
	if err != nil {
		return fmt.Errorf("InexRepo - CreateIncome - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	rows.Next()
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func (r InexRepo) ReadAllUsers(ctx context.Context) ([]domain.User, error) {
	rows, err := r.Pool.Query(ctx, sqlQueries.ReadUsers)
	if err != nil {
		return nil, fmt.Errorf("InexRepo - ReadAllUsers - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		if rows.Err() != nil {
			return nil, rows.Err()
		}

		var user domain.User
		err = rows.Scan(&user.ID)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
