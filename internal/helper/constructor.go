package helper

import (
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_code"
)

type Helper struct {
	MfaRecoveryCode mfa_recovery_code.MfaRecoveryCodesHelper
}

type Params struct {
	MfaRecoveryCode mfa_recovery_code.MfaRecoveryCodesHelper
}

func New(params Params) *Helper {
	helper := Helper{
		MfaRecoveryCode: params.MfaRecoveryCode,
	}
	return &helper
}
