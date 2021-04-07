package postgres

import (
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type ServiceProviderRepository struct {
	Tx transactor.Transactor
}
