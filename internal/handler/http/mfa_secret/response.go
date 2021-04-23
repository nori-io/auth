package mfa_secret

type MfaSecretResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Login   string `json:"login"`
	Issuer  string `json:"issuer"`
	Secret  string `json:"secret"`
}
