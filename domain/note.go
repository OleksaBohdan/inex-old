package domain

import (
	"github.com/gofrs/uuid"
)

type (
	Note struct {
		ID     uuid.UUID `json:"id"`
		UserId uuid.UUID `json:"user_id"`
		Text   string    `json:"text"`
	}
)
