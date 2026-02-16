package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
	"github.com/rigofekete/vhs-club-mvc/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/users", h.CreateUser)
	r.GET("/users/:id", h.GetUserByID)
	r.GET("/users", h.GetUsers)
	r.DELETE("/users", h.DeleteAllUsers)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser CreateUserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		_ = c.Error(apperror.WrapValidationError(err))
		return
	}

	createdUser, err := h.userService.CreateUser(c.Request.Context(), newUser.ToModel())
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, UserSingleResponse(createdUser))
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, UserSingleResponse(user))
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers(c.Request.Context())
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, UserListResponse(users))
}

func (h *UserHandler) DeleteAllUsers(c *gin.Context) {
	err := h.userService.DeleteAllUsers(c.Request.Context())
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
