package service

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationLogService interface {
	CreateAuthenticationLog(tx *gorm.DB, ctx context.Context, user *entity.User) error
}
