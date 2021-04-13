package settings

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

func newDisableMfaData(r *http.Request) (service.SecretData, error) {
	cSsid, err := r.Cookie("ssid")
	if err != nil {
		return service.SecretData{}, err
	}
	return service.SecretData{
		Ssid: cSsid.Value,
	}, nil
}
