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
}

func (h *TapeHandler) CreateTape(c *gin.Context) {
	var newTape model.Tape
	if err := c.ShouldBindJSON(&newTape); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := h.tapeService.Create(newTape)
	if created == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tape"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *TapeHandler) GetTapes(c *gin.Context) {
	tapes := h.tapeService.List()
	c.JSON(http.StatusOK, tapes)
}

func (h *TapeHandler) GetTapeByID(c *gin.Context) {
	id := c.Param("id")
	tape, found := h.tapeService.GetTapeByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "tape not found"})
		return
	}
	c.JSON(http.StatusOK, tape)
}

func (h *TapeHandler) UpdateTape(c *gin.Context) {
	id := c.Param("id")
	var update model.UpdatedTape
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tape, updated := h.tapeService.Update(id, update)
	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "error updating tape"})
		return
	}
	c.JSON(http.StatusOK, tape)
}

func (h *TapeHandler) DeleteTape(c *gin.Context) {
	id := c.Param("id")
	if deleted := h.tapeService.Delete(id); !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "tape not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
