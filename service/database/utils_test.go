package database_test

import (
	"testing"

	"github.com/magiconair/properties/assert"

	"github.com/nori-io/authentication/service/database"
)

func TestVerifyPassword_True(t *testing.T) {

	salt, _ := database.CreateSalt()

	cur, _ := database.Hash([]byte("pass"), salt)

	ok, _ := database.VerifyPassword([]byte("pass"), salt, cur)
	assert.Equal(t, ok, true)

}

func TestVerifyPassword_False(t *testing.T) {

	salt, _ := database.CreateSalt()

	cur, _ := database.Hash([]byte("pass"), salt)

	ok, _ := database.VerifyPassword([]byte("passFalse"), salt, cur)
	assert.Equal(t, ok, false)

}
