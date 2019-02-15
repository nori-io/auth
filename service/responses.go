package service

import "github.com/nori-io/auth/service/database"

//import "github.com/nori-io/noricms/service/database"

// SignUpResponse
type SignUpResponse struct {
	Id             uint64
	Name           string
	Email          string
	HttpStatusCode int
	Err            error
}

func (d *SignUpResponse) Error() error {
	return d.Err
}

func (d *SignUpResponse) StatusCode() int {
	return d.HttpStatusCode
}

// LogInResponse
type LogInResponse struct {
	Id             uint64
	Token          string
	User           database.UsersModel
	MFA            string
	HttpStatusCode int
	Err            error
}

func (d *LogInResponse) Error() error {
	return d.Err
}

func (d *LogInResponse) StatusCode() int {
	return d.HttpStatusCode
}

// LogOut Response
type LogOutResponse struct {
	HttpStatusCode int
	Err            error
}

func (d *LogOutResponse) Error() error {
	return d.Err
}

func (d *LogOutResponse) StatusCode() int {
	return d.HttpStatusCode
}
