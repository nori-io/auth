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

type SecretDataRequest struct {
	Secret string `json:"secret"`
}

func newPutSecretData(r *http.Request) (service.SecretData, error) {
	var body SecretDataRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return service.SecretData{}, err
	}

	cSsid, err := r.Cookie("ssid")
	if err != nil {
		return service.SecretData{}, err
	}
	return service.SecretData{
		Secret: body.Secret,
		Ssid:   cSsid.Value,
	}, nil
}
