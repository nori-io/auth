package settings

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-io/common/v5/pkg/domain/logger"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SettingsHandler struct {
	settingsService service.SettingsService
	cookieHelper    cookie.CookieHelper
	errorHelper     error2.ErrorHelper
	logger          logger.FieldLogger
}

type Params struct {
	SettingsService service.SettingsService
	CookieHelper    cookie.CookieHelper
	ErrorHelper     error2.ErrorHelper
	Logger          logger.FieldLogger
}

func New(params Params) *SettingsHandler {
	return &SettingsHandler{
		settingsService: params.SettingsService,
		cookieHelper:    params.CookieHelper,
		errorHelper:     params.ErrorHelper,
		logger:          params.Logger,
	}
}

func (h *SettingsHandler) DisableMfa(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	err = h.settingsService.DisableMfa(r.Context(), service.DisableMfaData{SessionKey: sessionId})
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, DisableMfaResponse{
		Success: true,
		Message: "mfa disabled",
	})
}

func (h *SettingsHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	data, err := newChangePasswordData(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.settingsService.ChangePassword(r.Context(), service.ChangePasswordData{
		SessionKey:  sessionId,
		PasswordOld: data.passwordOld,
		PasswordNew: data.passwordNew,
	})
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, PasswordChangeResponse{
		Success: true,
		Message: "password was changed",
	})
}

func (h *SettingsHandler) GetMfaStatus(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	enabled, err := h.settingsService.ReceiveMfaStatus(r.Context(), service.ReceiveMfaStatusData{SessionKey: sessionId})
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, MfaStatusResponse{
		Success: true,
		Message: "mfa status received",
		Status:  *enabled,
	})
}
