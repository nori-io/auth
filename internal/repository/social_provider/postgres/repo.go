package postgres

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type SocialProviderRepository struct {
	Tx transactor.Transactor
}

func (s SocialProviderRepository) Create(ctx context.Context, e *entity.SocialProvider) error {
	panic("implement me")
}

func (s SocialProviderRepository) Update(ctx context.Context, e *entity.SocialProvider) error {
	panic("implement me")
}

func (s SocialProviderRepository) Delete(ctx context.Context, ID uint64) error {
	panic("implement me")
}

func (s SocialProviderRepository) FindByID(ctx context.Context, ID uint64) (*entity.SocialProvider, error) {
	panic("implement me")
}

func (s SocialProviderRepository) FindByName(ctx context.Context, ID uint64) (*entity.SocialProvider, error) {
	panic("implement me")
}

func (s SocialProviderRepository) FindByFilter(ctx context.Context, filter repository.SocialProviderFilter) {
	panic("implement me")
}
