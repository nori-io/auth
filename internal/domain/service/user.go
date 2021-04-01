package service

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserService interface {
	CreateUser(tx *gorm.DB, ctx context.Context, data SignUpData) (*entity.User, error)
}
