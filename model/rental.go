package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Rental struct {
	ID         int32
	PublicID   uuid.UUID
	CreatedAt  time.Time
	UserID     int32
	TapeID     int32
	TapeTitle  string
	Username   string
	RentedAt   time.Time
	ReturnedAt sql.NullTime
}
