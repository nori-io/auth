package mfa_secret

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"

	"github.com/nori-io/common/v4/pkg/domain/logger"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaSecretHandler struct {
	mfaSecretService service.MfaSecretService
	logger           logger.FieldLogger
	cookieHelper     cookie.CookieHelper
	errorHelper      error2.ErrorHelper
}

type Params struct {
	MfaSecretService service.MfaSecretService
	Logger           logger.FieldLogger
	CookieHelper     cookie.CookieHelper
	ErrorHelper      error2.ErrorHelper
}

func New(params Params) *MfaSecretHandler {
	return &MfaSecretHandler{
		mfaSecretService: params.MfaSecretService,
		logger:           params.Logger,
	}
}

func (h *MfaSecretHandler) PutSecret(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	data, err := newPutSecretData(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//@todo
	login, issuer, err :=
		h.mfaSecretService.PutSecret(r.Context(), service.SecretData{
			Secret:     data.Secret,
			SessionKey: sessionId,
		})

	if (login == "") && (issuer == "") {
		h.logger.Error("%s", err)
		http.Error(w, "sign up error", http.StatusInternalServerError)
	}

	response.JSON(w, r, MfaSecretResponse{
		Success: true,
		Message: "secret",
		Login:   login,
		Issuer:  issuer,
		Secret:  "",
	})
}
