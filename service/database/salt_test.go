package database_test

import (
	"fmt"
	"testing"

	"github.com/nori-io/authorization/service/database"
)

func TestAuthenticate(t *testing.T) {

	salt, _ := database.Randbytes(8)
	fmt.Println("salt is", salt)
	fmt.Println("pass is", []byte("pass"))
	cur, _ := database.HashPassword([]byte("pass"), salt)
	fmt.Println("cur is", cur)

	if ok, _ := database.Authenticate([]byte("pass"), salt, cur); ok {
		fmt.Println("OK")
	}

}

func TestAuthenticate2(t *testing.T) {

	salt, _ := database.Randbytes(8)
	password, _ := database.HashPassword([]byte("pass"), salt)
	fmt.Println("salt is", salt)

	fmt.Println("cur is", password)

	var passwordArray = []byte{208, 148, 33, 206, 1, 82, 149, 86}
	var salt1 = []byte{29, 93, 90, 157, 70, 73, 81, 122}
	fmt.Println("Array salt", salt)
	fmt.Println("Array cur", passwordArray)

	result, err := database.Authenticate([]byte("pass"), salt1, passwordArray)
	fmt.Println(result, err)

	result, err = database.Authenticate([]byte("pass"), salt, password)
	fmt.Println(result, err)

}
