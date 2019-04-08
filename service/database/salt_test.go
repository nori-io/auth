package database_test

import (
	"fmt"
	"testing"

	"github.com/nori-io/authorization/service/database"
)

func TestAuthenticate(t *testing.T) {
	salt, _ := database.Randbytes(65)


	cur, _ := database.HashPassword([]byte("pass"), salt)

	if ok, _ := database.Authenticate([]byte("pass"), salt, cur); ok {
		fmt.Println("OK")
	}

	fmt.Println("No ok")
}



