package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rigofekete/vhs-club-mvc/service"
)

type RentalHandler struct {
	rentalService service.RentalService
}

func NewRentalHandler(s service.RentalService) *RentalHandler {
	return &RentalHandler{rentalService: s}
}

func (h *RentalHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/rentals/:id", h.CreateRental)
}

func (h *RentalHandler) CreateRental(c *gin.Context) {
	type parameters struct {
		UserID string `json:"user_id"`
	}

	tapeID := c.Param("id")
	var params parameters
	if err := c.ShouldBindJSON(&params); err != nil {
		_ = c.Error(err)
		return
	}
	createdRental, err := h.rentalService.RentTape(tapeID, params.UserID)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, createdRental)
}
