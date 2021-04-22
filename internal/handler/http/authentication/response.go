package authentication

import "time"

type SignInResponse struct {
	Success bool
	Message string
	MfaType string
}

type SignInMfaResponse struct {
	Success bool
	Message string
}

type SignInOutResponse struct {
	Success bool
	Message string
}

type SessionResponse struct {
	Success  bool
	Message  string
	Email    string
	Phone    string
	OpenedAt time.Time
}
