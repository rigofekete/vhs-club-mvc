package service

import "github.com/rigofekete/vhs-club-mvc/model"

type RentalService interface {
	Create(model.Rental) *model.Rental
}

type rentalService struct {
	repo repository.RentalRepository
}

func NewRentalService(r repository.RentalRepository) RentalService {
	return &rentalService{
		repo: r,
	}
}

func (s *rentalService) Create(rental model.Rental) *model.Rental {
	return s.repo.Save(rental)
}
