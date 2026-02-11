package service

import (
	"log"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type RentalService interface {
	Create(string, string) *model.Rental
}

type rentalService struct {
	repo repository.RentalRepository
}

func NewRentalService(r repository.RentalRepository) RentalService {
	return &rentalService{
		repo: r,
	}
}

func (s *rentalService) Create(tapeIDStr, userIDStr string) *model.Rental {
	tapeID, err := uuid.Parse(tapeIDStr)
	if err != nil {
		log.Printf("error parsing tape id string to uuid: %v", err)
		return nil
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		log.Printf("error parsing tape id string to uuid: %v", err)
		return nil
	}

	return s.repo.Save(tapeID, userID)
}
