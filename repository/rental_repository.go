package repository

import (
	"context"
	"sync"

	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type RentalRepository interface {
	Save(tape model.Rental) *model.Rental
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

func (r *rentalRepository) Save(rental model.Rental) *model.Rental {
	r.mu.Lock()
	defer r.mu.Unlock()
	rentalParams := database.CreateRentalParams{
		UserID: rental.UserID,
		TapeID: rental.TapeID,
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
