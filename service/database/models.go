package database

import (
	"time"
)

type UsersModel struct {
	Id             uint64
	Status_account string
	Type           string
	Created        time.Time
	Updated        time.Time
	Mfa_type       string
}

type AuthModel struct {
	Id              uint64
	UserId          uint64
	Phone           string
	Email           string
	Password        string
	Salt            string
	Created         time.Time
	Updated         time.Time
	IsEmailVerified bool
	IsPhoneVerified bool
}

type AuthProvidersModel struct {
	Provider        string
	ProviderUserKey string
	UserId          uint64
}

type AuthenticationHistoryModel struct {
	Id        uint64
	UserId    uint64
	LoggedIn  time.Time
	Meta      string
	LoggedOut time.Time
}

type UserMfaSecretModel struct {
	Id     uint64
	UserId uint64
	Secret string
}

type UserMfaPhoneModel struct {
	Id     uint64
	UserId uint64
	Phone  string
}

type UsersMfaCodeModel struct {
	Id     uint64
	UserId uint64
	code   string
}
