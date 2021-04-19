package social_provider

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/enum/social_provider_status"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (s SocialProviderService) GetAllActive(ctx context.Context) ([]entity.SocialProvider, error) {
	status := social_provider_status.Enabled

	providers, err := s.socialProviderRepository.FindByFilter(ctx, repository.SocialProviderFilter{
		Status: &status,
		Offset: 0,
		Limit:  0,
	})
	if err != nil {
		return nil, err
	}
	return providers, nil
}
