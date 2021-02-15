package mfa_secret

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/user/postgres"
)

func New(db *gorm.DB) repository.MfaSecret {
	return &postgres.MfaSecretRepository{Db: db}
}
