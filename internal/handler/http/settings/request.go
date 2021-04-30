package settings

import (
	"encoding/json"
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

func newDisableMfaData(r *http.Request) (service.DisableMfaData, error) {
	ssid, err := r.Cookie("ssid")
	if err != nil {
		return service.DisableMfaData{}, err
	}
	return service.DisableMfaData{
		SessionKey: ssid.Value,
	}, nil
}

type PasswordChangeRequest struct {
	passwordOld string `json:"password_old"`
	passwordNew string `json:"password_new"`
}

func newChangePasswordData(r *http.Request) (PasswordChangeRequest, error) {
	var body PasswordChangeRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return PasswordChangeRequest{}, err
	}
	return PasswordChangeRequest{
		passwordOld: body.passwordOld,
		passwordNew: body.passwordNew,
	}, nil
}
