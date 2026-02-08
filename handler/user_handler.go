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
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := h.userService.Create(newUser)
	c.JSON(http.StatusCreated, created)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.userService.List()
	c.JSON(http.StatusOK, users)
}
