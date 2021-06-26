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

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newLogInData(r *http.Request) (service.LogInData, error) {
	var body LogInRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return service.LogInData{}, err
	}
	return service.LogInData{
		Email:    body.Email,
		Password: body.Password,
	}, nil
}

type LogInMfaRequest struct {
	Code string `json:"code"`
}

func newLogInMfaData(r *http.Request) (service.LogInMfaData, error) {
	var body LogInMfaRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return service.LogInMfaData{}, err
	}
	return service.LogInMfaData{
		Code: body.Code,
	}, nil
}
