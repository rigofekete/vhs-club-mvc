package service

import (
	"strconv"

	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type RentalService interface {
	RentTape(string, string) (*model.Rental, error)
}

type rentalService struct {
	repo repository.RentalRepository
}

func NewRentalService(r repository.RentalRepository) RentalService {
	return &rentalService{
		repo: r,
	}
}

func (s *rentalService) RentTape(tapeIDStr, userIDStr string) (*model.Rental, error) {
	tapeID64, err := strconv.Atoi(tapeIDStr)
	if err != nil {
		return nil, err
	}
	userID64, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, err
	}

	return s.repo.Save(int32(tapeID64), int32(userID64))
}
