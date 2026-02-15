package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int32
	PublicID  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Email     string
}
