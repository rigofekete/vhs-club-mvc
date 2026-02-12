package model

import (
	"time"

	"github.com/google/uuid"
)

type Tape struct {
	ID        int32     `json:"id"`
	PublicID  uuid.UUID `json:"public_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Director  string    `json:"director"`
	Genre     string    `json:"genre"`
	Quantity  int32     `json:"quantity"`
	Price     float64   `json:"price"`
}

type UpdatedTape struct {
	Title    *string  `json:"title"`
	Director *string  `json:"director"`
	Genre    *string  `json:"genre"`
	Quantity *int32   `json:"quantity"`
	Price    *float64 `json:"price"`
}
