package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"

	"github.com/nori-plugins/authentication/pkg/enum/social_provider_status"
)

type SocialProviderRepository interface {
	Create(ctx context.Context, e *entity.SocialProvider) error
	Update(ctx context.Context, e *entity.SocialProvider) error
	Delete(ctx context.Context, ID uint64) error
	FindByID(ctx context.Context, ID uint64) (*entity.SocialProvider, error)
	FindByName(ctx context.Context, ID uint64) (*entity.SocialProvider, error)
	FindByFilter(ctx context.Context, filter SocialProviderFilter)
}

type SocialProviderFilter struct {
	Status *social_provider_status.Status
	Offset int
	Limit  int
}
