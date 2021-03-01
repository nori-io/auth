package mfa_recovery_codes

import "github.com/nori-plugins/authentication/internal/domain/entity"

type MfaRecoveryCodesHelper interface {
	GenerateRecoveryCodes(count uint8) []entity.MfaRecoveryCode
	GenerateRecoveryCode() entity.MfaRecoveryCode
}
