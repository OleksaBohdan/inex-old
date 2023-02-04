package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	User struct {
		ID           uuid.UUID
		Name         string
		Surname      string
		Email        string
		PasswordHash string
		RegisteredAt time.Time
		LastVisit    time.Time
	}
)
