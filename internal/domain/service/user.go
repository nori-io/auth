package service

import (
	"context"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"

	v "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserService interface {
	Create(ctx context.Context, data UserCreateData) (*entity.User, error)
	UpdatePassword(ctx context.Context, data UserUpdatePasswordData) error
	UpdateMfaStatus(ctx context.Context, data UserUpdateMfaStatusData) error
	GetByEmail(ctx context.Context, data GetByEmailData) (*entity.User, error)
	GetByID(ctx context.Context, data GetByIdData) (*entity.User, error)
	GetAll(ctx context.Context, data UserFilter) ([]entity.User, error)
}

type UserCreateData struct {
	Email    string
	Password string
}

func (d UserCreateData) Validate() error {
	return v.Errors{
		"email":    v.Validate(d.Email, v.Required, v.Length(3, 254), is.Email),
		"password": v.Validate(d.Password, v.Required),
	}.Filter()
}

type UserUpdateMfaStatusData struct {
	UserID  uint64
	MfaType mfa_type.MfaType
}

func (d UserUpdateMfaStatusData) Validate() error {
	return v.Errors{
		"user_ID":  v.Validate(d.UserID, v.Required),
		"mfa_type": v.Validate(d.MfaType, v.Required),
	}.Filter()
}

type UserUpdatePasswordData struct {
	UserID   uint64
	Password string
}

func (d UserUpdatePasswordData) Validate() error {
	return v.Errors{
		"email":    v.Validate(d.UserID, v.Required),
		"password": v.Validate(d.Password, v.Required),
	}.Filter()
}

type GetByIdData struct {
	Id uint64
}

func (d GetByIdData) Validate() error {
	return v.Errors{
		"id": v.Validate(d.Id, v.Required),
	}.Filter()
}

type GetByEmailData struct {
	Email string
}

func (d GetByEmailData) Validate() error {
	return v.Errors{
		"email": v.Validate(d.Email, v.Required, v.Length(3, 254), is.Email),
	}.Filter()
}

type UserFilter struct {
	EmailPattern *string
	PhonePattern *string
	UserStatus   *users_status.UserStatus
	Offset       int
	Limit        int
}
