package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newSignUpData(r *http.Request) (service.SignUpData, error) {
	var body SignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return service.SignUpData{}, err
	}
	return service.SignUpData{
		Email:    body.Email,
		Password: body.Password,
	}, nil
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newSignInData(r *http.Request) (service.SignInData, error) {
	var body SignInRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return service.SignInData{}, err
	}
	return service.SignInData{
		Email:    body.Email,
		Password: body.Password,
	}, nil
}

type SignInMfaRequest struct {
	Code string `json:"code"`
}

func newSignInMfaData(r *http.Request) (service.SignInMfaData, error) {
	var body SignInMfaRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return service.SignInMfaData{}, err
	}
	return service.SignInMfaData{
		Code: body.Code,
	}, nil
}
