package database_test

import (
	"fmt"
	"testing"

	"github.com/nori-io/authorization/service/database"
)

func TestAuthenticate(t *testing.T) {



	salt, _ := database.Randbytes(8)
	fmt.Println("salt is",salt)
    fmt.Println("pass is",[]byte("pass"))
	cur, _ := database.HashPassword([]byte("pass"), salt)
fmt.Println("cur is",cur)


	if ok, _ := database.Authenticate([]byte("pass"), salt, cur); ok {
		fmt.Println("OK")
	}

}

func TestAuthenticate2(t *testing.T) {



	salt, _ := database.Randbytes(8)
	fmt.Println("salt is",salt)
	fmt.Println([]byte("pass"))
	cur, _ := database.HashPassword([]byte("pass"), salt)
	fmt.Println("cur is",cur)


	var password=[]byte{63, 63, 63, 63, 89, 83, 74, 62}
	var salt1=[]byte{60, 22, 63, 63, 11, 63}

	if ok, _ := database.Authenticate([]byte("pass"), salt1, password); ok {
		fmt.Println("OK")
	}





	if ok, _ := database.Authenticate([]byte("pass"), salt, cur); ok {
		fmt.Println("OK")
	}

}

