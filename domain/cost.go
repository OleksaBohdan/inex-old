package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	Cost struct {
		ID        uuid.UUID
		UserId    uuid.UUID
		IncomeId  uuid.UUID
		Value     float64
		CreatedAt time.Time
	}
)
