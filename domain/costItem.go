package domain

import (
	"github.com/gofrs/uuid"
)

type (
	CostItem struct {
		ID     uuid.UUID `json:"id"`
		UserId uuid.UUID `json:"user_id"`
		Name   string    `json:"name"`
		Rang   int       `json:"rang"`
	}
)
