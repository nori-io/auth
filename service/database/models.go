package database

import (
	"time"
)

type AuthModel struct {
	Id_Auth             uint64
	UserId_Auth         uint64
	Phone_Auth           string
	Email_Auth           string
	Password_Auth        string
	Salt_Auth             string
	Created_Auth          time.Time
	Updated_Auth          time.Time
	IsEmailVerified_Auth  bool
	IsPhoneVerified_Auth  bool

	Id_Users       		 uint64
	Kind_Users  int64
	StatusId_Users       int64
	Type_Users           string
	Created_Users        time.Time
	Updated_Users        time.Time
	Mfa_type_Users          string


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
	Id            uint64
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
