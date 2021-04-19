package social_provider

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (s SocialProviderService) Get(ctx context.Context) ([]entity.SocialProvider, error) {
	providers, err := s.socialProviderRepository.FindByFilter(ctx, repository.SocialProviderFilter{
		Status: nil,
		Offset: 0,
		Limit:  0,
	})
	if err != nil {
		return nil, err
	}
	return providers, nil
}
