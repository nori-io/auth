package mfa_secret

import (
	"net/http"

	"github.com/nori-io/common/v4/pkg/domain/logger"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaSecretHandler struct {
	mfaSecretService service.MfaSecretService
	logger           logger.FieldLogger
}

type Params struct {
	MfaSecretService service.MfaSecretService
	Logger           logger.FieldLogger
}

func New(params Params) *MfaSecretHandler {
	return &MfaSecretHandler{
		mfaSecretService: params.MfaSecretService,
		logger:           params.Logger,
	}
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
		h.mfaSecretService.PutSecret(r.Context(), &service.SecretData{
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
