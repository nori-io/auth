package settings

import (
	s "github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SettingsService struct {
	sessionRepository repository.SessionRepository
	userRepository    repository.UserRepository
	session           s.Session
}

type Params struct {
	SessionRepository repository.SessionRepository
	UserRepository    repository.UserRepository
	Session           s.Session
}

func New(params Params) service.SettingsService {
	return &SettingsService{
		sessionRepository: params.SessionRepository,
		userRepository:    params.UserRepository,
		session:           params.Session,
	}
}
