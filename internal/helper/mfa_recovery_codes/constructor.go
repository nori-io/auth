package mfa_recovery_codes

import (
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_codes"
)

type mfaRecoveryCodesHelper struct {
}

func New() mfa_recovery_codes.MfaRecoveryCodesHelper {
	return &mfaRecoveryCodesHelper{}
}
