package settings

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SettingsService struct {
	userRepository repository.UserRepository
}

type Params struct {
	UserRepository repository.UserRepository
}

func New(params Params) service.SettingsService {
	return &SettingsService{
		userRepository: params.UserRepository,
	}
}
