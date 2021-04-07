package social_account

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/social_account/postgres"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

func New(tx transactor.Transactor) repository.SocialAccountRepository {
	return &postgres.SocialAccountRepository{Tx: tx}
}
