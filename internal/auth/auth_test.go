package auth_test

import (
	"testing"

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
	assert.NotEqual(t, passwordCorrect, hash)
}
