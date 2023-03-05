package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	Income struct {
		ID        uuid.UUID `json:"id"`
		UserId    uuid.UUID `json:"user_id"`
		IncomeId  uuid.UUID `json:"income_id"`
		Value     float64   `json:"value"`
		CreatedAt time.Time `json:"created_at"`
	}
)
