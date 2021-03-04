package mfa_recovery_code

import "github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_code"

type mfaRecoveryCodesHelper struct {
}

func New() mfa_recovery_code.MfaRecoveryCodesHelper {
	return &mfaRecoveryCodesHelper{}
}
