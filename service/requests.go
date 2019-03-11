package service

import (
	"github.com/asaskevich/govalidator"
	"github.com/cheebo/gorest"
)

// SignUp Request
type SignUpRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Type     string `json:"user_type"`
}

func (r SignUpRequest) Validate() error {



	_, err := govalidator.ValidateStruct(r)
	return rest.ValidateResponse(err)
}

// SignIn Request
type SignInRequest struct {
	Name     string
	Password string
}

func (r SignInRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return rest.ValidateResponse(err)
}

// SignOut Request
type SignOutRequest struct{}
