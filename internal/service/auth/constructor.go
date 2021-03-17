package auth

import (
	s "github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type AuthenticationService struct {
	AuthenticationHistoryRepository repository.AuthenticationHistoryRepository
	UserRepository                  repository.UserRepository
	SessionRepository               repository.SessionRepository
	Session                         s.Session
}

type Params struct {
	AuthenticationHistoryRepository repository.AuthenticationHistoryRepository
	SessionRepository               repository.SessionRepository
	UserRepository                  repository.UserRepository
	Session                         s.Session
}

func New(params Params) service.AuthenticationService {
	return &AuthenticationService{
		AuthenticationHistoryRepository: params.AuthenticationHistoryRepository,
		SessionRepository:               params.SessionRepository,
		UserRepository:                  params.UserRepository,
		Session:                         params.Session,
	}
}
