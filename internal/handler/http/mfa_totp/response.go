package mfa_totp

type MfaTotpResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Url     string `json:"url"`
}
