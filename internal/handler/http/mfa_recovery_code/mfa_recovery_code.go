package mfa_recovery_code

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type Handler struct {
	MfaRecoveryCodeService service.MfaRecoveryCodeService
}

func New(mfaRecoveryCodeService service.MfaRecoveryCodeService) *Handler {
	return &Handler{MfaRecoveryCodeService: mfaRecoveryCodeService}
}

func (h *Handler) GetMfaRecoveryCodes(w http.ResponseWriter, r *http.Request) {
	sessionIdContext := r.Context().Value("session_id")

	sessionId, _ := sessionIdContext.([]byte)

	if err := h.MfaRecoveryCodeService.GetMfaRecoveryCodes(r.Context(), &entity.Session{SessionKey: sessionId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//@todo path
	http.Redirect(w, r, "/", 0)
}
