package mfa_secret

type MfaSecretResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Email   string `json:"email"`
	Issuer  string `json:"issuer"`
	Secret  string `json:"secret"`
}
