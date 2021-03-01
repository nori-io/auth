package mfa_secret

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/mfa_secret/postgres"
)

func New(db *gorm.DB) repository.MfaSecretRepository {
	return &postgres.MfaSecretRepository{Db: db}
}
