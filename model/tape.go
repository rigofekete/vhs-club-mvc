package model

import (
	"time"

	"github.com/google/uuid"
)

type Tape struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Director  string    `json:"director"`
	Genre     string    `json:"genre"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
}
