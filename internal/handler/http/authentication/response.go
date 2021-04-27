package authentication

import "time"

type LogInResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	MfaType string `json:"mfa_type"`
}

type LogInMfaResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SessionResponse struct {
	Success  bool      `json:"success"`
	Message  string    `json:"message"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	OpenedAt time.Time `json:"opened_at"`
}
