package model

import (
	"time"

	"github.com/google/uuid"
)

type Tape struct {
	// ID        string    `json:"id"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Director  string    `json:"director"`
	Genre     string    `json:"genre"`
	Quantity  int       `json:"quantity"`
	// UserID    uuid.UUID `json:"user_id"`
	Price float64 `json:"price"`
}

// type User struct {
// 	ID string `json:"id"`
// }
