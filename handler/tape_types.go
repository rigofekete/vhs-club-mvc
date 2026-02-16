package handler

import (
	"time"

	"github.com/google/uuid"
)

type CreateTapeRequest struct {
	Title    string  `json:"title" binding:"required"`
	Director string  `json:"director" binding:"required"`
	Genre    string  `json:"genre" binding:"required"`
	Quantity int32   `json:"quantity" binding:"required,gt=0"`
	Price    float64 `json:"price" binding:"required,gt=0"`
}

type UpdateTapeRequest struct {
	Title    *string  `json:"title" binding:"omitempty,min=1,max=100"`
	Director *string  `json:"director" binding:"omitempty,min=1,max=50"`
	Genre    *string  `json:"genre" binding:"omitempty,min=1,max=50"`
	Quantity *int32   `json:"quantity" binding:"omitempty,gte=0"`
	Price    *float64 `json:"price" binding:"omitempty,gt=0"`
}

type TapeResponse struct {
	PublicID  uuid.UUID `json:"public_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Director  string    `json:"director"`
	Genre     string    `json:"genre"`
	Quantity  int32     `json:"quantity"`
	Price     float64   `json:"price"`
}
