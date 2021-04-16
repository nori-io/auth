package settings

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SettingsHandler struct {
	settingsService service.SettingsService
	logger          logger.FieldLogger
	cookieHelper    cookie.CookieHelper
	errorHelper     error2.ErrorHelper
}

type Params struct {
	SettingsService service.SettingsService
	Logger          logger.FieldLogger
	CookieHelper    cookie.CookieHelper
	ErrorHelper     error2.ErrorHelper
}

func New(params Params) *SettingsHandler {
	return &SettingsHandler{
		settingsService: params.SettingsService,
		logger:          params.Logger,
	}
}

func (h *SettingsHandler) DisableMfa(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	err = h.settingsService.DisableMfa(r.Context(), sessionId)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, http.StatusOK)
}
