package repository

import (
	"context"
	"sync"

	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type RentalRepository interface {
	Save(tapeID, userID int32) (*model.Rental, error)
	GetAll(ctx context.Context) ([]*model.Rental, error)
	GetActiveRentCount(ctx context.Context, tapeID int32) (int64, error)
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

	// TODO: Do this tape check in the service layer.....
	tape, err := r.DB.GetTapeByID(context.Background(), tapeID)
	if err != nil {
		return nil, apperror.ErrTapeNotFound
	}

	rentalParams := database.CreateRentalParams{
		UserID: userID,
		TapeID: tape.ID,
	}

	dbRental, err := r.DB.CreateRental(context.Background(), rentalParams)
	if err != nil {
		return nil, err
	}
	savedRental := &model.Rental{
		ID:         dbRental.ID,
		PublicID:   dbRental.PublicID,
		CreatedAt:  dbRental.CreatedAt,
		UserID:     dbRental.UserID,
		TapeID:     dbRental.TapeID,
		RentedAt:   dbRental.RentedAt,
		ReturnedAt: dbRental.ReturnedAt,
	}
	return savedRental, nil
}

func (r *rentalRepository) GetAll(ctx context.Context) ([]*model.Rental, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	dbRentals, err := r.DB.GetAllActiveRentals(ctx)
	if err != nil {
		return nil, err
	}

	rentals := make([]*model.Rental, 0)
	for _, rental := range dbRentals {
		r := &model.Rental{
			ID:        rental.ID,
			PublicID:  rental.PublicID,
			CreatedAt: rental.CreatedAt,
			UserID:    rental.UserID,
			TapeID:    rental.TapeID,
			RentedAt:  rental.RentedAt,
		}
		rentals = append(rentals, r)
	}
	return rentals, err
}

func (r *rentalRepository) GetActiveRentCount(ctx context.Context, tapeID int32) (int64, error) {
	count, err := r.DB.GetActiveRentalCountByTape(ctx, tapeID)
	if err != nil {
		return 0, err
	}
	return count, nil
}
