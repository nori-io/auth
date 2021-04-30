package postgres

import (
	"time"
	/*	"github.com/nori-plugins/authentication/pkg/enum/hash_algorithm"

		"github.com/nori-plugins/authentication/pkg/enum/mfa_type"
		"github.com/nori-plugins/authentication/pkg/enum/users_type"*/

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/pkg/enum/hash_algorithm"
	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"
	"github.com/nori-plugins/authentication/pkg/enum/users_status"
	"github.com/nori-plugins/authentication/pkg/enum/users_type"
)

type model struct {
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

func (m *model) convert() *entity.User {
	return &entity.User{
		ID:                     m.ID,
		Status:                 users_status.UserStatus(m.Status),
		UserType:               users_type.UserType(m.UserType),
		MfaType:                mfa_type.MfaType(m.MfaType),
		PhoneCountryCode:       m.PhoneCountryCode,
		PhoneNumber:            m.PhoneNumber,
		Email:                  m.Email,
		Password:               m.Password,
		Salt:                   m.Salt,
		HashAlgorithm:          hash_algorithm.HashAlgorithm(m.HashAlgorithm),
		IsEmailVerified:        m.IsEmailVerified,
		IsPhoneVerified:        m.IsPhoneVerified,
		EmailActivationCode:    m.EmailActivationCode,
		EmailActivationCodeTTL: m.EmailActivationCodeTTL,
		CreatedAt:              m.CreatedAt,
		UpdatedAt:              m.UpdatedAt,
	}
}

func newModel(e *entity.User) *model {
	return &model{
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
func (model) TableName() string {
	return "nori_authentication_users"
}
