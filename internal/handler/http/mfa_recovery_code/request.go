package mfa_recovery_code

import (
	"encoding/json"
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

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
