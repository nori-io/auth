package auth

import (
	s "github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type AuthenticationService struct {
	AuthenticationLogRepository repository.AuthenticationLogRepository
	UserRepository              repository.UserRepository
	SessionRepository           repository.SessionRepository
	MfaRecoveryCodeRepository   repository.MfaRecoveryCodeRepository
	Session                     s.Session
}

type Params struct {
	AuthenticationLogRepository repository.AuthenticationLogRepository
	MfaRecoveryCodeRepository   repository.MfaRecoveryCodeRepository
	SessionRepository           repository.SessionRepository
	UserRepository              repository.UserRepository
	Session                     s.Session
}

func New(params Params) service.AuthenticationService {
	return &AuthenticationService{
		AuthenticationLogRepository: params.AuthenticationLogRepository,
		MfaRecoveryCodeRepository:   params.MfaRecoveryCodeRepository,
		SessionRepository:           params.SessionRepository,
		UserRepository:              params.UserRepository,
		Session:                     params.Session,
	}
}
