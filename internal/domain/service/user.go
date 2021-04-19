package service

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserService interface {
	Create(ctx context.Context, data UserCreateData) (*entity.User, error)
	UpdatePassword(ctx context.Context, data UserUpdatePasswordData) error
	UpdateMfaStatus(ctx context.Context, data UserUpdateMfaStatusData) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, ID uint64) (*entity.User, error)
}

type UserCreateData struct {
	Email    string
	Password string
}

type UserUpdateMfaStatusData struct {
	UserID  uint64
	MfaType mfa_type.MfaType
}

type UserUpdatePasswordData struct {
	UserID   uint64
	Password string
}
