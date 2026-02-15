package handler

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=8,max=20"`
	Email    string `json:"email" binding:"required,email"`
}

type UserResponse struct {
	PublicID uuid.UUID `json:"public_id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}
