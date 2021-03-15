package service_provider

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/service_provider/postgres"
)

func New(db *gorm.DB) repository.ServiceProviderRepository {
	return &postgres.ServiceProviderRepository{Db: db}
}
