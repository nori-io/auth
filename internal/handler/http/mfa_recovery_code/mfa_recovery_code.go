package mfa_recovery_code

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"

	"github.com/nori-io/common/v4/pkg/domain/logger"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaRecoveryCodeHandler struct {
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	logger                 logger.FieldLogger
	cookieHelper           cookie.CookieHelper
	errorHelper            error2.ErrorHelper
}

type Params struct {
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	Logger                 logger.FieldLogger
	CookieHelper           cookie.CookieHelper
	ErrorHelper            error2.ErrorHelper
}

func New(params Params) *MfaRecoveryCodeHandler {
	return &MfaRecoveryCodeHandler{
		mfaRecoveryCodeService: params.MfaRecoveryCodeService,
		logger:                 params.Logger,
		cookieHelper:           params.CookieHelper,
		errorHelper:            params.ErrorHelper,
	}
}

func (h *MfaRecoveryCodeHandler) GetMfaRecoveryCodes(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	if _, err := h.mfaRecoveryCodeService.GetMfaRecoveryCodes(r.Context(), &entity.Session{SessionKey: []byte(sessionId)}); err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//@todo path
	http.Redirect(w, r, "/", 0)
}
