package auth

import (
	"github.com/jinzhu/gorm"
	s "github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type AuthenticationService struct {
	Config                      config.Config
	AuthenticationLogRepository repository.AuthenticationLogRepository
	UserRepository              repository.UserRepository
	SessionRepository           repository.SessionRepository
	MfaRecoveryCodeRepository   repository.MfaRecoveryCodeRepository
	DB                          *gorm.DB
	Session                     s.Session
}

type Params struct {
	Config                      config.Config
	AuthenticationLogRepository repository.AuthenticationLogRepository
	MfaRecoveryCodeRepository   repository.MfaRecoveryCodeRepository
	SessionRepository           repository.SessionRepository
	UserRepository              repository.UserRepository
	Session                     s.Session
	DB                          *gorm.DB
}

func New(params Params) service.AuthenticationService {
	return &AuthenticationService{
		Config:                      params.Config,
		AuthenticationLogRepository: params.AuthenticationLogRepository,
		MfaRecoveryCodeRepository:   params.MfaRecoveryCodeRepository,
		SessionRepository:           params.SessionRepository,
		UserRepository:              params.UserRepository,
		Session:                     params.Session,
		DB:                          params.DB,
	}
}
