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

func (h *MfaSecretHandler) GetSecret(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	email, issuer, err :=
		h.mfaSecretService.GetSecret(r.Context(), service.SecretData{
			SessionKey: sessionId,
		})

	if (email == "") && (issuer == "") {
		h.logger.Error("%s", err)
		http.Error(w, "sign up error", http.StatusInternalServerError)
	}

	response.JSON(w, r, MfaSecretResponse{
		Success: true,
		Message: "secret",
		Email:   email,
		Issuer:  issuer,
		Secret:  "",
	})
}
