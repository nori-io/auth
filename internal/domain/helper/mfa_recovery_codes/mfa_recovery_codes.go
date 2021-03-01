package mfa_recovery_codes

import "github.com/nori-plugins/authentication/internal/domain/entity"

type MfaRecoveryCodesHelper interface {
	GenerateRecoveryCodes(userID uint64, count int) ([]entity.MfaRecoveryCode, error)
	GenerateRecoveryCode() (entity.MfaRecoveryCode, error)
}
