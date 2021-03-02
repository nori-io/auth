package mfa_recovery_code

import (
	"context"

	s "github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/domain/repository"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type service struct {
	session                   s.Session
	userRepository            repository.UserRepository
	mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	mfaSecretRepository       repository.MfaSecretRepository
	config                    config
}

type config struct {
	Issuer string
}

func (srv *service) GetMfaRecoveryCodes(ctx context.Context, data *entity.Session) ([]entity.MfaRecoveryCode, error) {
	//@todo read count of symbols from config
	//@todo read pattenn from config
	//@todo read symbol sequence from config
	//@todo generating of specify sequence
	//@todo нужна ли максимальная длина, или указать всё в паттерне?
	err = srv.mfaRecoveryCodeRepository.Create(ctx, data.UserID, mfaRecoveryCode)

	return nil, nil
}
