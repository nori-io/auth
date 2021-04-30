package social_provider

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SocialProviderService struct {
	sessionService           service.SessionService
	socialProviderRepository repository.SocialProviderRepository
}

type Params struct {
	SessionService           service.SessionService
	SocialProviderRepository repository.SocialProviderRepository
}

func New(params Params) service.SocialProvider {
	return &SocialProviderService{
		sessionService:           params.SessionService,
		socialProviderRepository: params.SocialProviderRepository,
	}
}
