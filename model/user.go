package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int32     `json:"id"`
	PublicID  uuid.UUID `json:"public_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}
