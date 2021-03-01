package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/hash_algorithm"

	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"
	"github.com/nori-plugins/authentication/pkg/enum/users_type"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type User struct {
	ID                     uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	Status                 uint8     `gorm:"column:status; type:smallint; not null" `
	UserType               uint8     `gorm:"column:user_type; type:smallint; not null"`
	MfaType                uint8     `gorm:"column:mfa_type; type:smallint; null"`
	PhoneCountryCode       string    `gorm:"column:phone_country_code;uniqueIndex:idx_code_phone; type:VARCHAR(10)"`
	PhoneNumber            string    `gorm:"column:phone_number; uniqueIndex:idx_code_phone; type:VARCHAR(25)"`
	Email                  string    `gorm:"column:email; type:VARCHAR(254)"`
	Password               string    `gorm:"column:password; type:VARCHAR(32)"`
	Salt                   string    `gorm:"column:salt; type:VARCHAR(32)"`
	HashAlgorithm          uint8     `gorm:"column:hash_algorithm; type:smallint; not null"`
	IsEmailVerified        bool      `gorm:"column:is_email_verified; type:boolean; not null default false"`
	IsPhoneVerified        bool      `gorm:"column:is_phone_verified; type:boolean; not null default false"`
	EmailActivationCode    string    `gorm:"column:email_activation_code ; type:VARCHAR(254)"`
	EmailActivationCodeTTL time.Time `gorm:"column:email_activation_code_ttl; type: TIMESTAMP"`
	CreatedAt              time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
	UpdatedAt              time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}

func (u *User) Convert() *entity.User {
	return &entity.User{
		ID:                     u.ID,
		Status:                 users_status.UserStatus(u.Status),
		UserType:               users_type.UserType(u.UserType),
		MfaType:                mfa_type.MfaType(u.MfaType),
		PhoneCountryCode:       u.PhoneCountryCode,
		PhoneNumber:            u.PhoneNumber,
		Email:                  u.Email,
		Password:               u.Password,
		Salt:                   u.Salt,
		HashAlgorithm:          hash_algorithm.HashAlgorithm(u.HashAlgorithm),
		IsEmailVerified:        u.IsEmailVerified,
		IsPhoneVerified:        u.IsPhoneVerified,
		EmailActivationCode:    u.EmailActivationCode,
		EmailActivationCodeTTL: u.EmailActivationCodeTTL,
		CreatedAt:              u.CreatedAt,
		UpdatedAt:              u.UpdatedAt,
	}
}

func NewModel(e *entity.User) *User {
	return &User{
		ID:                     e.ID,
		Status:                 uint8(e.Status),
		UserType:               uint8(e.UserType),
		MfaType:                uint8(e.MfaType),
		PhoneCountryCode:       e.PhoneCountryCode,
		PhoneNumber:            e.PhoneNumber,
		Email:                  e.Email,
		Password:               e.Password,
		Salt:                   e.Salt,
		HashAlgorithm:          uint8(e.HashAlgorithm),
		IsEmailVerified:        e.IsEmailVerified,
		IsPhoneVerified:        e.IsPhoneVerified,
		EmailActivationCode:    e.EmailActivationCode,
		EmailActivationCodeTTL: e.EmailActivationCodeTTL,
		CreatedAt:              e.CreatedAt,
		UpdatedAt:              e.UpdatedAt,
	}
}

// TableName
func (User) TableName() string {
	return "users"
}
