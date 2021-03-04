package repository

import "github.com/nori-plugins/authentication/internal/domain/repository"

type Repository struct {
	UserRepository            repository.UserRepository
	MfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	MfaSecretRepository       repository.MfaSecretRepository
}

type Params struct {
	UserRepository            repository.UserRepository
	MfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	MfaSecretRepository       repository.MfaSecretRepository
}

func New(params Params) *Repository {
	repository := Repository{
		UserRepository:            params.UserRepository,
		MfaRecoveryCodeRepository: params.MfaRecoveryCodeRepository,
		MfaSecretRepository:       params.MfaSecretRepository,
	}
	return &repository
}
