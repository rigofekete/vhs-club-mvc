// Package handler provides controller logic
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/service"
)

type TapeHandler struct {
	tapeService service.TapeService
}

// Dependency Injection: TapeHandler depends on the TapeService abstraction
func NewTapeHandler(s service.TapeService) *TapeHandler {
	return &TapeHandler{tapeService: s}
}

func (h *TapeHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/tapes", h.CreateTape)
	r.GET("/tapes", h.GetTapes)
	r.GET("/tapes/:id", h.GetTapeByID)
	r.PATCH("/tapes/:id", h.UpdateTape)
	r.DELETE("/tapes/:id", h.DeleteTape)
	r.DELETE("/tapes", h.DeleteAllTapes)
}

func (h *TapeHandler) CreateTape(c *gin.Context) {
	var newTape model.Tape
	if err := c.ShouldBindJSON(&newTape); err != nil {
		_ = c.Error(err)
		return
	}
	createdTape, err := h.tapeService.CreateTape(newTape)
	if createdTape == nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tape"})
		return
	}
	c.JSON(http.StatusCreated, createdTape)
}

func (h *TapeHandler) GetTapes(c *gin.Context) {
	tapes, err := h.tapeService.ListTapes()
	if err != nil {
		_ = c.Error(err)
	}
	c.JSON(http.StatusOK, tapes)
}

func (h *TapeHandler) GetTapeByID(c *gin.Context) {
	id := c.Param("id")
	tape, err := h.tapeService.GetTapeByID(id)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, tape)
}

func (h *TapeHandler) UpdateTape(c *gin.Context) {
	id := c.Param("id")
	var update model.UpdatedTape
	if err := c.ShouldBindJSON(&update); err != nil {
		_ = c.Error(err)
		return
	}
	// TODO: Same names for methods in different layers. Check good practice.
	tape, err := h.tapeService.UpdateTape(id, update)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, tape)
}

func (h *TapeHandler) DeleteTape(c *gin.Context) {
	id := c.Param("id")
	if err := h.tapeService.DeleteTape(id); err != nil {
		_ = c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *TapeHandler) DeleteAllTapes(c *gin.Context) {
	if err := h.tapeService.DeleteAllTapes(); err != nil {
		_ = c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
