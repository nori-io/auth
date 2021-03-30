package mfa_recovery_code

import (
	"net/http"

	"github.com/nori-io/common/v4/pkg/domain/logger"

	s "github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/pkg/enum/session_status"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaRecoveryCodeHandler struct {
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	session                s.Session
	logger                 logger.FieldLogger
}

type Params struct {
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	session                s.Session
	logger                 logger.FieldLogger
}

func New(params Params) *MfaRecoveryCodeHandler {
	return &MfaRecoveryCodeHandler{
		mfaRecoveryCodeService: params.mfaRecoveryCodeService,
		session:                params.session,
		logger:                 params.logger,
	}
}

func (h *MfaRecoveryCodeHandler) GetMfaRecoveryCodes(w http.ResponseWriter, r *http.Request) {
	sessionId, err := r.Cookie("ssid")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	h.session.Get([]byte(sessionId.Value), session_status.Active)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	if _, err := h.mfaRecoveryCodeService.GetMfaRecoveryCodes(r.Context(), &entity.Session{SessionKey: []byte(sessionId.Value)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//@todo path
	http.Redirect(w, r, "/", 0)
}
