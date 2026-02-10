package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/service"
)

type RentalHandler struct {
	rentalService service.RentalService
}

func NewRentalHandler(s service.RentalService) *RentalHandler {
	return &RentalHandler{rentalService: s}
}

func (h *RentalHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/rentals", h.CreateRental)
}

func (h *RentalHandler) CreateRental(c *gin.Context) {
	var newRental model.Rental
	if err := c.ShouldBindJSON(&newRental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := h.rentalService.Create(newRental)
	if created == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid rental"})
		return
	}
	c.JSON(http.StatusCreated, created)
}
