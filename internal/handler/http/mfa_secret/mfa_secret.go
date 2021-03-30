package mfa_secret

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaSecretHandler struct {
	MfaSecretService service.MfaSecretService
}

func New(mfaSecretService service.MfaSecretService) *MfaSecretHandler {
	return &MfaSecretHandler{MfaSecretService: mfaSecretService}
}

func (h *MfaSecretHandler) PutSecret(w http.ResponseWriter, r *http.Request) {
	data, err := newPutSecretData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sessionIdContext := r.Context().Value("session_id")
	sessionId, _ := sessionIdContext.([]byte)

	if data.Ssid != sessionIdContext {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	sessionUserId := r.Context().Value("user_id").(uint64)
	//@todo
	login, issuer, err :=
		h.MfaSecretService.PutSecret(r.Context(), &service.SecretData{
			Secret: "",
			Ssid:   "",
		}, entity.Session{SessionKey: sessionId, UserID: sessionUserId})

	if (login == "") && (issuer == "") {
		http.Error(w, "sign up error", http.StatusInternalServerError)
	}

	response.JSON(w, r, MfaSecretResponse{
		Login:  login,
		Issuer: issuer,
	})
}
