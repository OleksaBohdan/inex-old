package domain

import (
	"github.com/gofrs/uuid"
)

type (
	IncomeItem struct {
		ID     uuid.UUID
		UserId uuid.UUID
		Name   string
	}
)
