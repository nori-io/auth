package mfa_recovery_codes

import (
	"crypto/rand"
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (h mfaRecoveryCodesHelper) GenerateRecoveryCodes(userID uint64, count int) ([]entity.MfaRecoveryCode, error) {
	var codes []entity.MfaRecoveryCode
	var mfaRecoveryCode entity.MfaRecoveryCode

	for i := 0; i < count; i++ {
		sid := make([]byte, 32)

		if _, err := rand.Read(sid); err != nil {
			return nil, err
		}

		mfaRecoveryCode = entity.MfaRecoveryCode{
			UserID:    userID,
			Code:      string(sid),
			CreatedAt: time.Now(),
		}
		codes = append(codes, mfaRecoveryCode)

	}
	return codes, nil
}

func (h mfaRecoveryCodesHelper) GenerateRecoveryCode() (entity.MfaRecoveryCode, error) {
	var mfaRecoveryCode entity.MfaRecoveryCode

	sid := make([]byte, 32)

	if _, err := rand.Read(sid); err != nil {
		return entity.MfaRecoveryCode{}, err
	}
	return mfaRecoveryCode, nil
}
