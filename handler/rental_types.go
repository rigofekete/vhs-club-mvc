package handler

import (
	"time"

	"github.com/google/uuid"
)

type CreateRentalRequest struct {
	UserPublicID string `json:"user_id" binding:"required"`
}

type ReturnRentalRequest struct {
	TapePublicID string `json:"tape_id" binding:"required"`
	UserPublicID string `json:"user_id" binding:"required"`
}

type RentalResponse struct {
	PublicID  uuid.UUID `json:"public_id"`
	TapeID    int32     `jsong:"tape_id"`
	UserID    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	RentedAt  time.Time `json:"rented_at"`
}
