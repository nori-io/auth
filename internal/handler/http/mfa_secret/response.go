package mfa_secret

type MfaSecretResponse struct {
	Success bool
	Message string
	Login   string
	Issuer  string
	Secret  string
}
