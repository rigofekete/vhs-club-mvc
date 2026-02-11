package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/users", h.CreateUser)
	r.GET("/users", h.GetUsers)
	// TODO: Add find User by id
	r.DELETE("/users", h.DeleteAllUsers)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO use DTO package to return DTO obj instead
	createdUser := h.userService.Create(newUser)
	// TODO: rename this created var, it is too ambiguous
	if createdUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user cannot be created"})
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.userService.List()
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) DeleteAllUsers(c *gin.Context) {
	h.userService.DeleteAll()
	c.Status(http.StatusNoContent)
}
