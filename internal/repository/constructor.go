package repository

import "github.com/nori-plugins/authentication/internal/domain/repository"

type Repository struct {
	AuthenticationLogRepository repository.AuthenticationLogRepository
	MfaRecoveryCodeRepository   repository.MfaRecoveryCodeRepository
	MfaTotpRepository           repository.MfaTotpRepository
	OneTimeTokenRepository      repository.OneTimeTokenRepository
	ServiceProviderRepository   repository.SocialProviderRepository
	SessionRepository           repository.SessionRepository
	SocialAccountRepository     repository.SocialAccountRepository
	UserRepository              repository.UserRepository
	ResetPasswordRepository     repository.ResetPasswordRepository
}

type Params struct {
	AuthenticationLogRepository repository.AuthenticationLogRepository
	MfaRecoveryCodeRepository   repository.MfaRecoveryCodeRepository
	MfaTotpRepository           repository.MfaTotpRepository
	OneTimeTokenRepository      repository.OneTimeTokenRepository
	ServiceProviderRepository   repository.SocialProviderRepository
	SessionRepository           repository.SessionRepository
	SocialAccountRepository     repository.SocialAccountRepository
	UserRepository              repository.UserRepository
	ResetPasswordRepository     repository.ResetPasswordRepository
}

func New(params Params) *Repository {
	repository := Repository{
		AuthenticationLogRepository: params.AuthenticationLogRepository,
		MfaRecoveryCodeRepository:   params.MfaRecoveryCodeRepository,
		MfaTotpRepository:           params.MfaTotpRepository,
		OneTimeTokenRepository:      params.OneTimeTokenRepository,
		ServiceProviderRepository:   params.ServiceProviderRepository,
		SessionRepository:           params.SessionRepository,
		SocialAccountRepository:     params.SocialAccountRepository,
		UserRepository:              params.UserRepository,
		ResetPasswordRepository:     params.ResetPasswordRepository,
	}
	return &repository
}
