package settings

import (
	"encoding/json"
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

func newDisableMfaData(r *http.Request) (service.SecretData, error) {
	ssid, err := r.Cookie("ssid")
	if err != nil {
		return service.SecretData{}, err
	}
	return service.SecretData{
		SessionKey: ssid.Value,
	}, nil
}

type ChangePasswordData struct {
	passwordOld string `json:"password_old"`
	passwordNew string `json:"password_new"`
}

func newChangePasswordData(r *http.Request) (ChangePasswordData, error) {
	var body ChangePasswordData

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return ChangePasswordData{}, err
	}
	return ChangePasswordData{
		passwordOld: body.passwordOld,
		passwordNew: body.passwordNew,
	}, nil
}
