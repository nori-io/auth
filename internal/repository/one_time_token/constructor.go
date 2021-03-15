package one_time_token

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/one_time_token/postgres"
)

func New(db *gorm.DB) repository.OneTimeTokenRepository {
	return &postgres.OneTimeTokenRepository{Db: db}
}
