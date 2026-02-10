package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Rental struct {
	ID         uuid.UUID    `json:"id"`
	CreatedAt  time.Time    `json:"created_at"`
	UserID     uuid.UUID    `json:"user_id"`
	TapeID     uuid.UUID    `json:"tape_id"`
	RentedAt   time.Time    `json:"rented_at"`
	ReturnedAt sql.NullTime `json:"returned_at"`
}
