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
