package database

import (
	"time"
)

type UsersModel struct {
	Id       uint64
	Status_account int64
	Type     string
	Created  time.Time
	Updated  time.Time
	Mfa_type string
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

type AuthenticationHistoryModel struct {
	Id        int64
	UserId    int64
	LoggedIn  time.Time
	Meta      string
	LoggedOut time.Time
	Secret    string
}

type AuthProvidersModel struct {
	Provider        string
	ProviderUserKey string
}



type UserStatusModel struct {
	Id   int64
	Name string
}
