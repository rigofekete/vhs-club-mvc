package apperror

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code    int
	Message string
}

// Sentinel Errors
var (
	// General
	ErrBadRequest = errors.New("bad request")
	// User
	ErrUserNotFound   = errors.New("user not found")
	ErrUserValidation = errors.New("invalid user fields")
	// Tape
	ErrTapeValidation = errors.New("invalid tape fields")
	ErrTapeNotFound   = errors.New("tape not found")
)

func mapErrorToAppError(err error) *AppError {
	switch {
	case errors.Is(err, ErrBadRequest):
		return &AppError{Code: 400, Message: "Bad request"}
	case errors.Is(err, ErrUserNotFound):
		return &AppError{Code: 404, Message: "User not found"}
	case errors.Is(err, ErrUserValidation):
		return &AppError{Code: 422, Message: "Invalid user fields"}
	case errors.Is(err, ErrTapeValidation):
		return &AppError{Code: 422, Message: "Invalid tape fields"}
	case errors.Is(err, ErrTapeNotFound):
		return &AppError{Code: 404, Message: "Tape not found"}
	default:
		return &AppError{Code: 500, Message: "Internal server error"}
	}
}

// Middleware factory function
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}
		appErr := mapErrorToAppError(err.Err)
		c.JSON(appErr.Code, gin.H{
			"error": appErr.Message,
		})
	}
}
