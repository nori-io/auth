package mfa_recovery_code

import (
	"context"
	"time"

	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_codes"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	service2 "github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeService struct {
	mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	mfaRecoveryCodeHelper     mfa_recovery_codes.MfaRecoveryCodesHelper
	config                    Config
}

type Config struct {
	MfaRecoveryCodeCount int
}

func New(mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository, mfaRecoveryCodeHelper mfa_recovery_codes.MfaRecoveryCodesHelper, config Config) service2.MfaRecoveryCodeService {
	return &MfaRecoveryCodeService{
		mfaRecoveryCodeRepository: mfaRecoveryCodeRepository,
		mfaRecoveryCodeHelper:     mfaRecoveryCodeHelper,
		config:                    config,
	}
}

func (srv *MfaRecoveryCodeService) GetMfaRecoveryCodes(ctx context.Context, data *entity.Session) ([]entity.MfaRecoveryCode, error) {
	//@todo read count of symbols from config
	//@todo read pattenn from config
	//@todo read symbol sequence from config
	//@todo generating of specify sequence
	//@todo нужна ли максимальная длина, или указать всё в паттерне?
	var mfaRecoveryCodes []entity.MfaRecoveryCode
	mfa_recovery_codes, err := srv.mfaRecoveryCodeHelper.Generate(srv.config.MfaRecoveryCodeCount)
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
	err = srv.mfaRecoveryCodeRepository.Create(ctx, mfaRecoveryCodes)

	return nil, nil
}
