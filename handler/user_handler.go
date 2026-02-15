package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/service"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r *CreateUserRequest) ToModel() *model.User {
	return &model.User{
		Name:  r.Name,
		Email: r.Email,
	}
}

type UserResponse struct {
	PublicID uuid.UUID `json:"public_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
}

func UserSingleResponse(user *model.User) *UserResponse {
	return &UserResponse{
		PublicID: user.PublicID,
		Name:     user.Name,
		Email:    user.Email,
	}
}

func UserListResponse(user []*model.User) []*UserResponse {
	userList := make([]*UserResponse, len(user))
	for i, user := range user {
		userList[i] = UserSingleResponse(user)
	}
	return userList
}

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
	var newUser CreateUserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		_ = c.Error(err)
		return
	}

	createdUser, err := h.userService.CreateUser(newUser.ToModel())
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, UserSingleResponse(createdUser))
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.ListUsers()
	if err != nil {
		_ = c.Error(err)
	}
	c.JSON(http.StatusOK, UserListResponse(users))
}

func (h *UserHandler) DeleteAllUsers(c *gin.Context) {
	err := h.userService.DeleteAllUsers()
	if err != nil {
		_ = c.Error(err)
	}
	c.Status(http.StatusNoContent)
}
