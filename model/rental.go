package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Rental struct {
	ID         int32        `json:"id"`
	PublicID   uuid.UUID    `json:"public_id"`
	CreatedAt  time.Time    `json:"created_at"`
	UserID     int32        `json:"user_id"`
	TapeID     int32        `json:"tape_id"`
	RentedAt   time.Time    `json:"rented_at"`
	ReturnedAt sql.NullTime `json:"returned_at"`
}
