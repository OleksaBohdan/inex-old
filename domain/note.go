package domain

import (
	"github.com/gofrs/uuid"
)

type (
	Note struct {
		ID     uuid.UUID
		UserId uuid.UUID
		Text   string
	}
)
