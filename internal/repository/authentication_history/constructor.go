package authentication_history

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/authentication_history/postgres"
)

func New(db *gorm.DB) repository.AuthenticationHistoryRepository {
	return &postgres.AuthenticationHistoryRepository{Db: db}
}
