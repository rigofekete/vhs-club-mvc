package repository

import (
	"context"
	"sync"

	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type RentalRepository interface {
	Save(ctx context.Context, tapeID, userID int32) (*model.Rental, error)
	GetAllActive(ctx context.Context) ([]*model.Rental, error)
	GetActiveRentCountByTape(ctx context.Context, tapeID int32) (*int64, error)
	GetActiveRentCountByUser(ctx context.Context, userID int32) (*int64, error)
	DeleteAllRentals(ctx context.Context) error
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

func (r *rentalRepository) Save(ctx context.Context, tapeID, userID int32) (*model.Rental, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	rentalParams := database.CreateRentalParams{
		UserID: userID,
		TapeID: tapeID,
	}

	dbRental, err := r.DB.CreateRental(ctx, rentalParams)
	if err != nil {
		return nil, err
	}

	savedRental := &model.Rental{
		ID:         dbRental.ID,
		PublicID:   dbRental.PublicID,
		CreatedAt:  dbRental.CreatedAt,
		UserID:     dbRental.UserID,
		TapeID:     dbRental.TapeID,
		TapeTitle:  dbRental.Title,
		Username:   dbRental.Username,
		RentedAt:   dbRental.RentedAt,
		ReturnedAt: dbRental.ReturnedAt,
	}
	return savedRental, nil
}

func (r *rentalRepository) GetAllActive(ctx context.Context) ([]*model.Rental, error) {
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
			TapeTitle: rental.Title,
			Username:  rental.Username,
			RentedAt:  rental.RentedAt,
		}
		rentals = append(rentals, r)
	}
	return rentals, err
}

func (r *rentalRepository) GetActiveRentCountByTape(ctx context.Context, tapeID int32) (*int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	count, err := r.DB.GetActiveRentalCountByTape(ctx, tapeID)
	if err != nil {
		return nil, err
	}
	return &count, nil
}

func (r *rentalRepository) GetActiveRentCountByUser(ctx context.Context, userID int32) (*int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	count, err := r.DB.GetActiveRentalCountByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &count, nil
}

func (r *rentalRepository) DeleteAllRentals(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.DB.DeleteAllRentals(ctx); err != nil {
		return err
	}
	return nil
}
