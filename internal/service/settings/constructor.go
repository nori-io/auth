package settings

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SettingsService struct {
	sessionRepository repository.SessionRepository
	userRepository    repository.UserRepository
}

type Params struct {
	SessionRepository repository.SessionRepository
	UserRepository    repository.UserRepository
}

func New(params Params) service.SettingsService {
	return &SettingsService{
		sessionRepository: params.SessionRepository,
		userRepository:    params.UserRepository,
	}
}
