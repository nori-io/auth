package service

import (
	"github.com/asaskevich/govalidator"
	"github.com/cheebo/gorest"
)

// SignUp Request
type SignUpRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Type        string `json:"user_type"`
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
