package authentication_log

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/authentication_log/postgres"
)

func New(db *gorm.DB) repository.AuthenticationLogRepository {
	return &postgres.AuthenticationLogRepository{Db: db}
}
