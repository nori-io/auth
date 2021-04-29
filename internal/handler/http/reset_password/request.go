package reset_password

import (
	"encoding/json"
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

func newRequestResetPasswordEmailData(r *http.Request) (service.RequestResetPasswordEmailData, error) {
	var body ResetPasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return service.RequestResetPasswordEmailData{}, err
	}
	return service.RequestResetPasswordEmailData{
		Email: body.Email,
	}, nil
}

type ResetPasswordSetRequest struct {
	token    string `json:"token"`
	password string `json:"password"`
}

func NewSetNewPasswordByResetPasswordEmailToken(r *http.Request) (service.SetNewPasswordByResetPasswordEmailTokenData, error) {
	var body ResetPasswordSetRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return service.SetNewPasswordByResetPasswordEmailTokenData{}, err
	}
	return service.SetNewPasswordByResetPasswordEmailTokenData{}, nil
}
