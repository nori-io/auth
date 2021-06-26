package mfa_totp

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/helper"

	"github.com/nori-io/common/v5/pkg/domain/logger"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaTotpHandler struct {
	mfaTotpService service.MfaTotpService
	cookieHelper   helper.CookieHelper
	errorHelper    helper.ErrorHelper
	logger         logger.FieldLogger
}

type Params struct {
	MfaTotpService service.MfaTotpService
	CookieHelper   helper.CookieHelper
	ErrorHelper    helper.ErrorHelper
	Logger         logger.FieldLogger
}

func New(params Params) *MfaTotpHandler {
	return &MfaTotpHandler{
		mfaTotpService: params.MfaTotpService,
		cookieHelper:   params.CookieHelper,
		errorHelper:    params.ErrorHelper,
		logger:         params.Logger,
	}
}

func (h *MfaTotpHandler) GetUrl(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	url, err := h.mfaTotpService.GetUrl(r.Context(), service.MfaGetUrlData{
		SessionKey: sessionId,
	})
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, MfaTotpResponse{
		Success: true,
		Message: "",
		Url:     url,
	})
}
