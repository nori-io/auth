package mfa_secret

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	service "github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaSecretService struct {
	MfaSecretRepository repository.MfaSecretRepository
	UserRepository      repository.UserRepository
	Config              config.Config
}

type Params struct {
	MfaSecretRepository repository.MfaSecretRepository
	UserRepository      repository.UserRepository
	Config              config.Config
}

func New(params Params) service.MfaSecretService {
	return &MfaSecretService{MfaSecretRepository: params.MfaSecretRepository, UserRepository: params.UserRepository}
}
