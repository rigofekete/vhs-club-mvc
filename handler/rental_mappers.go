package handler

import "github.com/rigofekete/vhs-club-mvc/model"

// func (r *CreateRentalRequest) ToModel() *model.Rental {
// 	return &model.Rental{
// 		UserID: r.UserPublicID,
// 	}
// }
//

func RentalSingleResponse(rental *model.Rental) *RentalResponse {
	return &RentalResponse{
		PublicID:  rental.PublicID,
		CreatedAt: rental.CreatedAt,
		RentedAt:  rental.RentedAt,
		UserID:    rental.UserID,
		TapeID:    rental.TapeID,
	}
}

func RentalListResponse(rentals []*model.Rental) []*RentalResponse {
	rentalList := make([]*RentalResponse, len(rentals))
	for i, rental := range rentals {
		rentalList[i] = RentalSingleResponse(rental)
	}
	return rentalList
}
