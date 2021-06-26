package helper

import (
	"github.com/nori-plugins/authentication/internal/domain/helper"
	"github.com/nori-plugins/authentication/internal/helper/goth_provider"
)

type Helper struct {
	CookieHelper          helper.CookieHelper
	ErrorHelper           helper.ErrorHelper
	GothProviderHelper    goth_provider.GothProviderHelper
	MfaRecoveryCodeHelper helper.MfaRecoveryCodesHelper
	MfaTotpHelper         helper.MfaTotpHelper
	SecurityHelper        helper.SecurityHelper
}

type Params struct {
	CookieHelper          helper.CookieHelper
	ErrorHelper           helper.ErrorHelper
	GothProviderHelper    goth_provider.GothProviderHelper
	MfaRecoveryCodeHelper helper.MfaRecoveryCodesHelper
	MfaTotpHelper         helper.MfaTotpHelper
	SecurityHelper        helper.SecurityHelper
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
