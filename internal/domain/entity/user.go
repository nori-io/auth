package entity

import (
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/hash_algorithm"

	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"
	"github.com/nori-plugins/authentication/pkg/enum/users_type"
)

type User struct {
	ID                  uint64
	Status              users_status.UserStatus
	UserType            users_type.UserType
	MfaType             mfa_type.MfaType
	PhoneCountryCode    string
	PhoneNumber         string
	Email               string
	Password            string
	Salt                string
	HashAlgorithm       hash_algorithm.HashAlgorithm
	IsEmailVerified     bool
	IsPhoneVerified     bool
	EmailActivationCode string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
