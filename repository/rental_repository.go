package repository

import (
	"context"
	"database/sql"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type RentalRepository interface {
	Save(tapeID, userID uuid.UUID) *model.Rental
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

func (r *rentalRepository) Save(tapeID, userID uuid.UUID) *model.Rental {
	r.mu.Lock()
	defer r.mu.Unlock()

	tape, err := r.DB.GetTape(context.Background(), tapeID)
	if err != nil {
		// TODO: decide how to handle such errors back to the controller layer
		log.Printf("error: requested tape does not exist in the database: %v", err)
		return nil
	}

	rentalParams := database.CreateRentalParams{
		UserID:     userID,
		TapeID:     tape.ID,
		ReturnedAt: sql.NullTime{Valid: false},
	}

	dbRental, err := r.DB.CreateRental(context.Background(), rentalParams)
	if err != nil {
		// TODO: Should we return the err together with the object pointer?
		return nil
	}
	savedRental := &model.Rental{
		ID:         dbRental.ID,
		CreatedAt:  dbRental.CreatedAt,
		UserID:     dbRental.UserID,
		TapeID:     dbRental.TapeID,
		RentedAt:   dbRental.RentedAt,
		ReturnedAt: dbRental.ReturnedAt,
	}
	return savedRental
}
