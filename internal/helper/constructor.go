package helper

import (
	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
)

type Helper struct {
	CookieHelper          cookie.CookieHelper
	ErrorHelper           error2.ErrorHelper
	MfaRecoveryCodeHelper mfa_recovery_code.MfaRecoveryCodesHelper
	SecurityHelper        security.SecurityHelper
}

type Params struct {
	CookieHelper    cookie.CookieHelper
	ErrorHelper     error2.ErrorHelper
	MfaRecoveryCode mfa_recovery_code.MfaRecoveryCodesHelper
	SecurityHelper  security.SecurityHelper
}

func New(params Params) *Helper {
	helper := Helper{
		CookieHelper:          params.CookieHelper,
		ErrorHelper:           params.ErrorHelper,
		MfaRecoveryCodeHelper: params.MfaRecoveryCode,
		SecurityHelper:        params.SecurityHelper,
	}
	return &helper
}
