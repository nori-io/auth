package mfa_recovery_code

import (
	"context"
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"
)

//@todo как передать сюда всю сессию? скорее, нужно извлечь пользовательский userID из контекста
//@todo что мы будем хранить в контексте?
func (srv *MfaRecoveryCodeService) GetMfaRecoveryCodes(ctx context.Context, data *entity.Session) ([]*entity.MfaRecoveryCode, error) {
	//@todo будет ли использоваться паттерн?
	//@todo нужна ли максимальная длина, или указать всё в паттерне?
	//@todo указать ограничение на максимальную длину, связанную с базой данных?

	var mfaRecoveryCodes []*entity.MfaRecoveryCode
	mfa_recovery_codes, err := srv.mfaRecoveryCodeHelper.Generate()
	if err != nil {
		return nil, err
	}
	for _, v := range mfa_recovery_codes {
		mfaRecoveryCodes = append(mfaRecoveryCodes, &entity.MfaRecoveryCode{
			ID:        0,
			UserID:    data.UserID,
			Code:      v,
			CreatedAt: time.Now(),
		})
	}

	if err := srv.transactor.Transact(ctx, func(tx context.Context) error {
		if err = srv.mfaRecoveryCodeRepository.DeleteMfaRecoveryCodes(ctx, data.UserID); err != nil {
			return err
		}
		if err = srv.mfaRecoveryCodeRepository.Create(ctx, mfaRecoveryCodes); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return mfaRecoveryCodes, nil
}

func (srv *MfaRecoveryCodeService) GetByUserId(ctx context.Context, userID uint64, code string) (*entity.MfaRecoveryCode, error) {
	//@todo проверить кэш и отп
	mfaRecoveryCode, err := srv.mfaRecoveryCodeRepository.FindByUserID(ctx, userID, code)
	if err != nil {
		return nil, err
	}
	if mfaRecoveryCode == nil {
		return nil, errors2.MfaRecoveryCodeNotFound
	}
	return mfaRecoveryCode, nil
}

func (srv *MfaRecoveryCodeService) Apply(ctx context.Context, userID uint64, code string) error {
	if err := srv.mfaRecoveryCodeRepository.DeleteMfaRecoveryCode(ctx, userID, code); err != nil {
		return err
	}
	return nil
}
