package mfa_recovery_codes

import "github.com/nori-plugins/authentication/internal/domain/entity"

func (h helper) GenerateRecoveryCodes(count uint8) []entity.MfaRecoveryCode {
	return nil
}

func (h helper) GenerateRecoveryCode() entity.MfaRecoveryCode {
	return entity.MfaRecoveryCode{}
}
