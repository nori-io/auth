package helper

import (
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
)

type Helper struct {
	MfaRecoveryCodeHelper mfa_recovery_code.MfaRecoveryCodesHelper
	SecurityHelper        security.SecurityHelper
}

type Params struct {
	MfaRecoveryCode mfa_recovery_code.MfaRecoveryCodesHelper
	SecurityHelper  security.SecurityHelper
}

func New(params Params) *Helper {
	helper := Helper{
		MfaRecoveryCodeHelper: params.MfaRecoveryCode,
		SecurityHelper:        params.SecurityHelper,
	}
	return &helper
}
