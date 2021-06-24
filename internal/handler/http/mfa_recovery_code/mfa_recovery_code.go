package mfa_recovery_code

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/entity"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"

	"github.com/nori-io/common/v5/pkg/domain/logger"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaRecoveryCodeHandler struct {
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	cookieHelper           cookie.CookieHelper
	errorHelper            error2.ErrorHelper
	logger                 logger.FieldLogger
}

type Params struct {
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	CookieHelper           cookie.CookieHelper
	ErrorHelper            error2.ErrorHelper
	Logger                 logger.FieldLogger
}

func New(params Params) *MfaRecoveryCodeHandler {
	return &MfaRecoveryCodeHandler{
		mfaRecoveryCodeService: params.MfaRecoveryCodeService,
		cookieHelper:           params.CookieHelper,
		errorHelper:            params.ErrorHelper,
		logger:                 params.Logger,
	}
}

func (h *MfaRecoveryCodeHandler) GetMfaRecoveryCodes(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	mfaCodes, err := h.mfaRecoveryCodeService.GetMfaRecoveryCodes(r.Context(), service.GetMfaRecoveryCodes{SessionKey: sessionId})
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	codes := convertAll(mfaCodes)

	response.JSON(w, r, MfaRecoveryCodesResponse{
		success: true,
		message: "mfa recovery codes generated",
		codes:   codes,
	})
}

func convertAll(entities []*entity.MfaRecoveryCode) []string {
	codes := make([]string, 0)
	for _, v := range entities {
		codes = append(codes, convert(*v))
	}
	return codes
}

func convert(e entity.MfaRecoveryCode) string {
	return e.Code
}
