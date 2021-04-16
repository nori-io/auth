package social_provider

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/social_provider/postgres"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

func New(tx transactor.Transactor) repository.SocialProviderRepository {
	return &postgres.SocialProviderRepository{Tx: tx}
}
