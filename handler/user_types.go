package handler

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=4,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type CreateUserBatchRequest struct {
	// dive validator needed to iterate each object of the slice and apply the CreateUserRequest binded validations
	Users []CreateUserRequest `json:"users" binding:"required,dive"`
}

type UserResponse struct {
	PublicID uuid.UUID `json:"public_id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

type UserBatchResponse struct {
	Users         []UserResponse `json:"tapes"`
	AlreadyExists int32          `json:"already_exists"`
}
