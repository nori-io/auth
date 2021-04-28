package mfa_totp

type MfaTotpHelper interface {
	Generate(email string) (url string, secret string, err error)
}
