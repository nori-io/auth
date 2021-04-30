package repository

import "github.com/nori-plugins/authentication/internal/domain/repository"

type Repository struct {
	MfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	MfaTotpRepository         repository.MfaTotpRepository
	OneTimeTokenRepository    repository.OneTimeTokenRepository
	ResetPasswordRepository   repository.ResetPasswordRepository
	ServiceProviderRepository repository.SocialProviderRepository
	SessionRepository         repository.SessionRepository
	SocialAccountRepository   repository.SocialAccountRepository
	UserRepository            repository.UserRepository
	UserLogRepository         repository.UserLogRepository
}

type Params struct {
	MfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	MfaTotpRepository         repository.MfaTotpRepository
	OneTimeTokenRepository    repository.OneTimeTokenRepository
	ResetPasswordRepository   repository.ResetPasswordRepository
	ServiceProviderRepository repository.SocialProviderRepository
	SessionRepository         repository.SessionRepository
	SocialAccountRepository   repository.SocialAccountRepository
	UserRepository            repository.UserRepository
	UserLogRepository         repository.UserLogRepository
}

func New(params Params) *Repository {
	repository := Repository{
		MfaRecoveryCodeRepository: params.MfaRecoveryCodeRepository,
		MfaTotpRepository:         params.MfaTotpRepository,
		OneTimeTokenRepository:    params.OneTimeTokenRepository,
		ResetPasswordRepository:   params.ResetPasswordRepository,
		ServiceProviderRepository: params.ServiceProviderRepository,
		SessionRepository:         params.SessionRepository,
		SocialAccountRepository:   params.SocialAccountRepository,
		UserRepository:            params.UserRepository,
		UserLogRepository:         params.UserLogRepository,
	}
	return &repository
}
