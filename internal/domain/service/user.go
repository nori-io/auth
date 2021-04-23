package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	v "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserService interface {
	Create(ctx context.Context, data UserCreateData) (*entity.User, error)
	UpdatePassword(ctx context.Context, data UserUpdatePasswordData) error
	UpdateMfaStatus(ctx context.Context, data UserUpdateMfaStatusData) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, data GetByIdData) (*entity.User, error)
	GetAll(ctx context.Context, filter repository.UserFilter) ([]entity.User, error)
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

type GetByIdData struct {
	Id uint64
}

func (d GetByIdData) Validate() error {
	return v.Errors{
		"id": v.Validate(d.Id, v.Required),
	}.Filter()
}
