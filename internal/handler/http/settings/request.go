package settings

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

func newDisableMfaData(r *http.Request) (service.DisableMfaData, error) {
	cSsid, err := r.Cookie("ssid")
	if err != nil {
		return service.DisableMfaData{}, err
	}
	return service.SecretData{
		Secret: body.Secret,
		Ssid:   cSsid.Value,
	}, nil
}
