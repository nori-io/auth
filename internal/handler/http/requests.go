package http

import (
	"github.com/asaskevich/govalidator"
	rest "github.com/cheebo/gorest"
)

// SignUp Request
type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r SignUpRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return rest.ValidateResponse(err)
}

// LogIn Request
type SignInRequest struct {
	Email    string
	Password string
}

func (r SignInRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return rest.ValidateResponse(err)
}

// LogOut Request
type SignOutRequest struct{}
