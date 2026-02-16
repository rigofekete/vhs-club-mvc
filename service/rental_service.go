package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type RentalService interface {
	RentTape(ctx context.Context, tapeID string, userID string) (*model.Rental, error)
	GetAllRentals(ctx context.Context) ([]*model.Rental, error)
}

type rentalService struct {
	tapeRepo   repository.TapeRepository
	userRepo   repository.UserRepository
	rentalRepo repository.RentalRepository
}

func NewRentalService(r repository.RentalRepository, t repository.TapeRepository, u repository.UserRepository) RentalService {
	return &rentalService{
		rentalRepo: r,
		tapeRepo:   t,
		userRepo:   u,
	}
}

func (s *rentalService) RentTape(ctx context.Context, tapePublicID, userPublicID string) (*model.Rental, error) {
	tapeUUID, err := uuid.Parse(tapePublicID)
	if err != nil {
		return nil, err
	}
	userUUID, err := uuid.Parse(userPublicID)
	if err != nil {
		return nil, err
	}

	// TODO: We probably don't need this GetIDFromPublicID but GetTape/UserByPublicID instead, since we need to check more fields
	tapeID, err := s.tapeRepo.GetIDFromPublicID(ctx, tapeUUID)
	if err != nil {
		return nil, fmt.Errorf("could not find tape from given public id: %w: %v", apperror.ErrTapeNotFound, err)
	}

	userID, err := s.userRepo.GetIDFromPublicID(ctx, userUUID)
	if err != nil {
		return nil, fmt.Errorf("could not find user from given public id: %w: %v", apperror.ErrUserNotFound, err)
	}

	tape, err := s.tapeRepo.GetByID(ctx, tapeID)
	if err != nil {
		return nil, err
	}

	count, err := s.rentalRepo.GetActiveRentCount(ctx, tapeID)
	if err != nil {
		return nil, err
	}

	if int32(count) >= tape.Quantity {
		return nil, apperror.ErrTapeUnavailable
	}

	return s.rentalRepo.Save(tapeID, userID)
}

func (s *rentalService) GetAllRentals(ctx context.Context) ([]*model.Rental, error) {
	return s.rentalRepo.GetAll(ctx)
}
