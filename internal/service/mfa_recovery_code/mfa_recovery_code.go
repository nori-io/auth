package mfa_recovery_code

import (
	"context"
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

//@todo как передать сюда всю сессию? скорее, нужно извлечь пользовательский userID из контекста
//@todo что мы будем хранить в контексте?
func (srv *MfaRecoveryCodeService) GetMfaRecoveryCodes(ctx context.Context, data *entity.Session) ([]entity.MfaRecoveryCode, error) {
	//@todo будет ли использоваться паттерн?
	//@todo нужна ли максимальная длина, или указать всё в паттерне?
	//@todo указать ограничение на максимальную длину, связанную с базой данных?

	var mfaRecoveryCodes []entity.MfaRecoveryCode
	mfa_recovery_codes, err := srv.mfaRecoveryCodeHelper.Generate()
	if err != nil {
		return nil, err
	}
	for _, v := range mfa_recovery_codes {
		mfaRecoveryCodes = append(mfaRecoveryCodes, entity.MfaRecoveryCode{
			ID:        0,
			UserID:    data.UserID,
			Code:      v,
			CreatedAt: time.Now(),
		})
	}
	tx := srv.db.Begin()
	if err = srv.mfaRecoveryCodeRepository.DeleteMfaRecoveryCodes(tx, ctx, data.UserID); err != nil {
		return nil, err
	}
	if err = srv.mfaRecoveryCodeRepository.Create(tx, ctx, mfaRecoveryCodes); err != nil {
		return nil, err
	}

	tx.Commit()

	return mfaRecoveryCodes, nil
}
