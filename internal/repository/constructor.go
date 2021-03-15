package repository

import "github.com/nori-plugins/authentication/internal/domain/repository"

type Repository struct {
	AuthenticationHistoryRepository repository.AuthenticationHistoryRepository
	MfaRecoveryCodeRepository       repository.MfaRecoveryCodeRepository
	MfaSecretRepository             repository.MfaSecretRepository
	OneTimeTokenRepository          repository.OneTimeTokenRepository
	ServiceProviderRepository       repository.ServiceProviderRepository
	SessionRepository               repository.SessionRepository
	SocialAccountRepository         repository.SocialAccountRepository
	UserRepository                  repository.UserRepository
}

type Params struct {
	AuthenticationHistoryRepository repository.AuthenticationHistoryRepository
	MfaRecoveryCodeRepository       repository.MfaRecoveryCodeRepository
	MfaSecretRepository             repository.MfaSecretRepository
	OneTimeTokenRepository          repository.OneTimeTokenRepository
	ServiceProviderRepository       repository.ServiceProviderRepository
	SessionRepository               repository.SessionRepository
	SocialAccountRepository         repository.SocialAccountRepository
	UserRepository                  repository.UserRepository
}

func New(params Params) *Repository {
	repository := Repository{
		AuthenticationHistoryRepository: params.AuthenticationHistoryRepository,
		MfaRecoveryCodeRepository:       params.MfaRecoveryCodeRepository,
		MfaSecretRepository:             params.MfaSecretRepository,
		OneTimeTokenRepository:          params.OneTimeTokenRepository,
		ServiceProviderRepository:       params.ServiceProviderRepository,
		SessionRepository:               params.SessionRepository,
		SocialAccountRepository:         params.SocialAccountRepository,
		UserRepository:                  params.UserRepository,
	}
	return &repository
}
