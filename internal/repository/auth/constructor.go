package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/auth/postgres"
)

func New(db *gorm.DB) repository.AuthRepository {
	return &postgres.AuthRepository{Db: db}
}
