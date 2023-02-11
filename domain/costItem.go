package domain

import (
	"github.com/gofrs/uuid"
)

type (
	CostItem struct {
		ID     uuid.UUID
		UserId uuid.UUID
		Name   string
	}
)
