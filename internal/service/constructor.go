package service

import (
	s "github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/pkg"
)

type Service struct {
	session                   s.Session
	userRepository            repository.UserRepository
	mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	mfaSecretRepository       repository.MfaSecretRepository
	config                    config
}

type config struct {
	Issuer string
}

func New(sessionInstance s.Session,
	userRepository repository.UserRepository,
	mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository,
	mfaSecretRepository repository.MfaSecretRepository,
	config config) pkg.Authentication {
	return &Service{
		config:                    config,
		session:                   sessionInstance,
		userRepository:            userRepository,
		mfaRecoveryCodeRepository: mfaRecoveryCodeRepository,
		mfaSecretRepository:       mfaSecretRepository,
	}
}
