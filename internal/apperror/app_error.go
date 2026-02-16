package apperror

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Code    int
	Message string
	Fields  map[string]string `json:"fields,omitempty"`
}

// Sentinel Errors
var (
	// General
	ErrBadRequest = errors.New("bad request")
	// User
	ErrUserNotFound   = errors.New("user not found")
	ErrUserValidation = errors.New("invalid user fields")
	// Tape
	ErrTapeValidation    = errors.New("invalid tape fields")
	ErrTapeNotFound      = errors.New("tape not found")
	ErrTapeUpdateRequest = errors.New("bad update tape request")
	ErrTapeUnavailable   = errors.New("unavailable tape")
)

type ValidationError struct {
	Fields map[string]string
}

// ValidationError type implements the Error interface, so we can pass it as an error in the handler layer, c.ShouldBindJSON error check
func (e ValidationError) Error() string {
	return "input validation failed"
}

func WrapValidationError(err error) error {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	fields := make(map[string]string)

	for _, fieldError := range validationErrors {
		field := fieldError.Field()

		switch fieldError.Tag() {
		case "required":
			fields[field] = "This field is required"
		case "email":
			fields[field] = "Must be a valid email address"
		case "alphanum":
			fields[field] = "Must be formed of only letters and/or numbers"
		case "min":
			fields[field] = "Must be at least " + fieldError.Param() + " characters"
		case "max":
			fields[field] = "Must be at most " + fieldError.Param() + " characters"
		default:
			fields[field] = "Invalid value"
		}
	}
	return ValidationError{Fields: fields}
}

func mapErrorToAppError(err error) *AppError {
	// Check if incoming error is a ValidationError
	var validationErr ValidationError
	if errors.As(err, &validationErr) {
		return &AppError{
			Code:    400,
			Message: "Input validation failed",
			Fields:  validationErr.Fields,
		}
	}

	// Check sentinel errors
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
	case errors.Is(err, ErrTapeUpdateRequest):
		return &AppError{Code: 404, Message: "Tape update request needs at least 1 non nil value"}
	case errors.Is(err, ErrTapeUnavailable):
		return &AppError{Code: 404, Message: "Sorry, all the tapes for this movie are currently rented out"}
	default:
		return &AppError{Code: 500, Message: "Internal server error"}
	}
}

// TODO: Should this be moved to its own middleware package ?
// Middleware factory function
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}
		appErr := mapErrorToAppError(err.Err)

		response := gin.H{
			"error": appErr.Message,
		}

		if appErr.Fields != nil {
			response["fields"] = appErr.Fields
		}

		c.JSON(appErr.Code, response)
	}
}
