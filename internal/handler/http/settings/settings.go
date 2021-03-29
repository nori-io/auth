package settings

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SettingsHandler struct {
	settingsService service.SettingsService
	logger          logger.FieldLogger
}

type Params struct {
	SettingsService service.SettingsService
	Logger          logger.FieldLogger
}

func New(params Params) *SettingsHandler {
	return &SettingsHandler{
		settingsService: params.SettingsService,
		logger:          params.Logger,
	}
}

func (h *SettingsHandler) DisableMfa(w http.ResponseWriter, r *http.Request) {
	sessionIdContext := r.Context().Value("session_id")

	sessionId, _ := sessionIdContext.(string)

	err := h.settingsService.DisableMfa(r.Context(), sessionId)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, http.StatusOK)
}
