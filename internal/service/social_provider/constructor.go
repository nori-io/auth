package social_provider

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SocialProviderService struct {
	socialProviderRepository repository.SocialProviderRepository
}

type Params struct {
	SocialProviderRepository repository.SocialProviderRepository
}

func New(params Params) service.SocialProvider {
	return &SocialProviderService{
		socialProviderRepository: params.SocialProviderRepository,
	}
}
