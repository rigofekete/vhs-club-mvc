package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
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
	var req CreateRentalRequest

	tapeID := c.Param("id")
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(apperror.WrapValidationError(err))
		return
	}
	createdRental, err := h.rentalService.RentTape(c.Request.Context(), tapeID, req.UserPublicID)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, createdRental)
}
