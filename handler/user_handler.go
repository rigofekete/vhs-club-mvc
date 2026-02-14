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
		_ = c.Error(err)
		return
	}

	// TODO: use DTO package to return DTO obj instead
	createdUser, err := h.userService.CreateUser(newUser)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.ListUsers()
	if err != nil {
		_ = c.Error(err)
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) DeleteAllUsers(c *gin.Context) {
	err := h.userService.DeleteAllUsers()
	if err != nil {
		_ = c.Error(err)
	}
	c.Status(http.StatusNoContent)
}
