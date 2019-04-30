package database_test

import (
	"testing"

	"github.com/magiconair/properties/assert"

	"github.com/nori-io/authorization/service/database"
)

func TestAuthenticate(t *testing.T) {

	salt, _ := database.Randbytes(8)

	cur, _ := database.Hash([]byte("pass"), salt)

	ok, _ := database.Authenticate([]byte("pass"), salt, cur)
	assert.Equal(t, ok, true)

}

func TestAuthenticate2(t *testing.T) {

	salt, _ := database.Randbytes(8)
	password, _ := database.HashPassword([]byte("pass"), salt)

	var passwordArray = []byte{208, 148, 33, 206, 1, 82, 149, 86}
	var salt1 = []byte{29, 93, 90, 157, 70, 73, 81, 122}

	result, _ := database.Authenticate([]byte("pass"), salt1, passwordArray)
	assert.Equal(t, result, true)

	result, _ = database.Authenticate([]byte("pass"), salt, password)
	assert.Equal(t, result, true)

}
