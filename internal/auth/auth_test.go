package auth_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/internal/auth"
	"github.com/stretchr/testify/assert"
)

func Test_CheckPasswordHash_Success(t *testing.T) {
	password := "guybrushThreepwood1991!"
	hash, _ := auth.HashPassword(password)

	valid, err := auth.CheckPasswordHash(password, hash)

	assert.Nil(t, err)
	assert.Equal(t, valid, true)
	assert.NotEqual(t, password, hash)
}

func Test_CheckPasswordHash_WrongPassword(t *testing.T) {
	passwordCorrect := "guybrushThreepwood1991!"
	passwordWrong := "LeChuck91!"
	hash, _ := auth.HashPassword(passwordCorrect)

	valid, err := auth.CheckPasswordHash(passwordWrong, hash)

	assert.Nil(t, err)
	assert.Equal(t, valid, false)
}

func Test_ValidateJWT(t *testing.T) {
	userID := uuid.New()
	jwtTokenString, _ := auth.MakeJWT(userID, "admin", "secret", time.Hour)

	validatedID, role, err := auth.ValidateJWT(jwtTokenString, "secret")

	assert.Nil(t, err)
	assert.Equal(t, role, "admin")
	assert.Equal(t, userID, validatedID)
}

func Test_ValidateJWT_InvalidRole(t *testing.T) {
	userID := uuid.New()
	_, _ = auth.MakeJWT(userID, "user", "secret", time.Hour)

	_, _, err := auth.ValidateJWT("someRandomHack", "secret")

	assert.Error(t, err)
}
