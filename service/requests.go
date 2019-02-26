package service

import (
	"github.com/asaskevich/govalidator"
	"github.com/cheebo/gorest"
)

// SignUp Request
type SignUpRequest struct {
	Email    string
	Password string
}

func (r SignUpRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return rest.ValidateResponse(err)
}

// LogIn Request
type LoginRequest struct {
	Email    string
	Password string
}

func (r LoginRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return rest.ValidateResponse(err)
}

// LogOut Request
type LogoutRequest struct{}
