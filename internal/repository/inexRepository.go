package repository

import (
	"context"
	"inex/main/domain"
)

type (
	InexRepository struct{}
)

func (r InexRepository) CreateUser(ctx context.Context, user domain.User) {

}
