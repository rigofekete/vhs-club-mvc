package repository

import (
	"context"
	"database/sql"
	"errors"
	"sync"

	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type RentalRepository interface {
	Save(tapeID, userID int32) (*model.Rental, error)
}

type rentalRepository struct {
	mu sync.Mutex
	DB *database.Queries
}

func NewRentalRepository() RentalRepository {
	return &rentalRepository{
		DB: config.AppConfig.DB,
	}
}

func (r *rentalRepository) Save(tapeID, userID int32) (*model.Rental, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tape, err := r.DB.GetTape(context.Background(), tapeID)
	if err != nil {
		// TODO: should we return the raw error coming from sql?
		return nil, errors.New("requested tape does not exist in the database")
	}

	rentalParams := database.CreateRentalParams{
		UserID:     userID,
		TapeID:     tape.ID,
		ReturnedAt: sql.NullTime{Valid: false},
	}

	dbRental, err := r.DB.CreateRental(context.Background(), rentalParams)
	if err != nil {
		return nil, err
	}
	savedRental := &model.Rental{
		ID:         dbRental.ID,
		PublicID:   dbRental.PublicID.UUID,
		CreatedAt:  dbRental.CreatedAt,
		UserID:     dbRental.UserID,
		TapeID:     dbRental.TapeID,
		RentedAt:   dbRental.RentedAt,
		ReturnedAt: dbRental.ReturnedAt,
	}
	return savedRental, nil
}
