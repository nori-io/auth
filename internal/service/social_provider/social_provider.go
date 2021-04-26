package social_provider

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/service"

	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"

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

func (srv *SocialProviderService) GetByName(ctx context.Context, data service.GetByNameData) (*entity.SocialProvider, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	provider, err := srv.socialProviderRepository.FindByName(ctx, data.Name)
	if err != nil {
		return nil, err
	}
	if provider == nil {
		return nil, errors2.SocialProviderNotFound
	}

	return provider, nil
}