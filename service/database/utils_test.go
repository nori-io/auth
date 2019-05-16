package database_test

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"

	"github.com/nori-io/authentication/service/database"
)

func TestVerifyPassword(t *testing.T) {

	salt, _ := database.CreateSalt()
	fmt.Println("Salt is", salt)

	cur, _ := database.Hash([]byte("pass"), salt)

	fmt.Println("Cur is", cur)

	ok, _ := database.VerifyPassword([]byte("pass"), salt, cur)
	assert.Equal(t, ok, true)

}


