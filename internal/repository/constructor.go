package repository

import "github.com/nori-plugins/authentication/internal/domain/repository"

type Repository struct {
	AuthenticationLogRepository repository.AuthenticationLogRepository
	MfaRecoveryCodeRepository   repository.MfaRecoveryCodeRepository
	MfaSecretRepository         repository.MfaSecretRepository
	OneTimeTokenRepository      repository.OneTimeTokenRepository
	ServiceProviderRepository   repository.SocialProviderRepository
	SessionRepository           repository.SessionRepository
	SocialAccountRepository     repository.SocialAccountRepository
	UserRepository              repository.UserRepository
}

type Params struct {
	AuthenticationLogRepository repository.AuthenticationLogRepository
	MfaRecoveryCodeRepository   repository.MfaRecoveryCodeRepository
	MfaSecretRepository         repository.MfaSecretRepository
	OneTimeTokenRepository      repository.OneTimeTokenRepository
	ServiceProviderRepository   repository.SocialProviderRepository
	SessionRepository           repository.SessionRepository
	SocialAccountRepository     repository.SocialAccountRepository
	UserRepository              repository.UserRepository
}

func New(params Params) *Repository {
	repository := Repository{
		AuthenticationLogRepository: params.AuthenticationLogRepository,
		MfaRecoveryCodeRepository:   params.MfaRecoveryCodeRepository,
		MfaSecretRepository:         params.MfaSecretRepository,
		OneTimeTokenRepository:      params.OneTimeTokenRepository,
		ServiceProviderRepository:   params.ServiceProviderRepository,
		SessionRepository:           params.SessionRepository,
		SocialAccountRepository:     params.SocialAccountRepository,
		UserRepository:              params.UserRepository,
	}
	return &repository
}
