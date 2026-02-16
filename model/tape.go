package model

import (
	"time"

	"github.com/google/uuid"
)

type Tape struct {
	ID        int32
	PublicID  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Director  string
	Genre     string
	Quantity  int32
	Price     float64
}

type UpdateTape struct {
	ID       int32
	Title    *string
	Director *string
	Genre    *string
	Quantity *int32
	Price    *float64
}
