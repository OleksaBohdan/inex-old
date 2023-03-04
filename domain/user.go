package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	User struct {
		ID           uuid.UUID `json:"id"`
		Name         string    `json:"name"`
		Surname      string    `json:"surname"`
		Email        string    `json:"email"`
		PasswordHash string    `json:"password_hash"`
		RegisteredAt time.Time `json:"registered_at"`
		LastVisit    time.Time `json:"last_visit"`
	}
)
