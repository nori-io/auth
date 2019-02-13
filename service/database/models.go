package database

import (
	"time"
)

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
	UserId          int64
}

type UsersModel struct {
	Id            int64
	ProfileTypeId int64
	StatusId      int64
	Kind          string
	Created       time.Time
	Updated       time.Time
	Email         string
}

type UserStatusModel struct {
	Id   int64
	Name string
}
