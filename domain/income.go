package domain

import (
	"github.com/gofrs/uuid"
)

type (
	Income struct {
		ID       uuid.UUID
		UserId   uuid.UUID
		IncomeId uuid.UUID
		Value    float64
	}
)
