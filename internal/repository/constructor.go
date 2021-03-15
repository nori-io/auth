package repository

import "github.com/nori-plugins/authentication/internal/domain/repository"

type Repository struct {
	MfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	MfaSecretRepository       repository.MfaSecretRepository
	UserRepository            repository.UserRepository
}

type Params struct {
	MfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	MfaSecretRepository       repository.MfaSecretRepository
	UserRepository            repository.UserRepository
}

func New(params Params) *Repository {
	repository := Repository{
		MfaRecoveryCodeRepository: params.MfaRecoveryCodeRepository,
		MfaSecretRepository:       params.MfaSecretRepository,
		UserRepository:            params.UserRepository,
	}
	return &repository
}
