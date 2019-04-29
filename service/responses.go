package service

import (
	"github.com/nori-io/authentication/service/database"
)

// SignUpResponse
type SignUpResponse struct {
	Id               uint64
	Name             string
	Email            string
	PhoneCountryCode string
	PhoneNumber      string
	HttpStatusCode   int
	Err              error
}

func (d *SignUpResponse) Error() error {
	return d.Err
}

func (d *SignUpResponse) StatusCode() int {
	return d.HttpStatusCode
}

// SignInResponse
type SignInResponse struct {
	Id             uint64
	Token          string
	User           database.AuthModel
	MFA            string
	HttpStatusCode int
	Err            error
}

func (d *SignInResponse) Error() error {
	return d.Err
}

func (d *SignInResponse) StatusCode() int {
	return d.HttpStatusCode
}

// SignOut Response
type SignOutResponse struct {
	HttpStatusCode int
	Err            error
}

func (d *SignOutResponse) Error() error {
	return d.Err
}

func (d *SignOutResponse) StatusCode() int {
	return d.HttpStatusCode
}

type RecoveryCodesResponse struct {
	Codes          []string
	HttpStatusCode int
	Err            error
}

func (d *RecoveryCodesResponse) Error() error {
	return d.Err
}

func (d *RecoveryCodesResponse) StatusCode() int {
	return d.HttpStatusCode
}
