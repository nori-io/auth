package service_provider

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/service_provider/postgres"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

func New(tx transactor.Transactor) repository.ServiceProviderRepository {
	return &postgres.ServiceProviderRepository{Tx: tx}
}
