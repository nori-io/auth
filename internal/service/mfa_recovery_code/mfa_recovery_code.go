package mfa_recovery_code

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	service2 "github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeService struct {
	mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
}

func New(mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository) service2.MfaRecoveryCodeService {
	return &MfaRecoveryCodeService{mfaRecoveryCodeRepository: mfaRecoveryCodeRepository}
}

func (srv *MfaRecoveryCodeService) GetMfaRecoveryCodes(ctx context.Context, data *entity.Session) ([]entity.MfaRecoveryCode, error) {
	//@todo read count of symbols from config
	//@todo read pattenn from config
	//@todo read symbol sequence from config
	//@todo generating of specify sequence
	//@todo нужна ли максимальная длина, или указать всё в паттерне?
	err = srv.mfaRecoveryCodeRepository.Create(ctx, data.UserID, mfaRecoveryCode)

	return nil, nil
}
