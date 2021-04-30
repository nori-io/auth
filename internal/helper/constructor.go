package helper

import (
	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_totp"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/helper/goth_provider"
)

type Helper struct {
	CookieHelper          cookie.CookieHelper
	ErrorHelper           error2.ErrorHelper
	GothProviderHelper    goth_provider.GothProviderHelper
	MfaRecoveryCodeHelper mfa_recovery_code.MfaRecoveryCodesHelper
	MfaTotpHelper         mfa_totp.MfaTotpHelper
	SecurityHelper        security.SecurityHelper
}

type Params struct {
	CookieHelper          cookie.CookieHelper
	ErrorHelper           error2.ErrorHelper
	GothProviderHelper    goth_provider.GothProviderHelper
	MfaRecoveryCodeHelper mfa_recovery_code.MfaRecoveryCodesHelper
	MfaTotpHelper         mfa_totp.MfaTotpHelper
	SecurityHelper        security.SecurityHelper
}

func New(params Params) *Helper {
	helper := Helper{
		CookieHelper:          params.CookieHelper,
		ErrorHelper:           params.ErrorHelper,
		GothProviderHelper:    params.GothProviderHelper,
		MfaRecoveryCodeHelper: params.MfaRecoveryCodeHelper,
		MfaTotpHelper:         params.MfaTotpHelper,
		SecurityHelper:        params.SecurityHelper,
	}
	return &helper
}
