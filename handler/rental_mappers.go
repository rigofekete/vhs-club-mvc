package handler

import "github.com/rigofekete/vhs-club-mvc/model"

func RentalSingleResponse(rental *model.Rental) *RentalResponse {
	return &RentalResponse{
		PublicID:  rental.PublicID,
		CreatedAt: rental.CreatedAt,
		RentedAt:  rental.RentedAt,
		UserID:    rental.UserID,
		TapeID:    rental.TapeID,
		TapeTitle: rental.TapeTitle,
		Username:  rental.Username,
	}
}

func RentalListResponse(rentals []*model.Rental) []*RentalResponse {
	rentalList := make([]*RentalResponse, len(rentals))
	for i, rental := range rentals {
		rentalList[i] = RentalSingleResponse(rental)
	}
	return rentalList
}
