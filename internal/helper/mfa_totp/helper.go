package mfa_totp

import (
	"github.com/pquerna/otp/totp"
)

func (h totpHelper) Generate(email string) (url string, secret string, err error) {
	opts := totp.GenerateOpts{
		Issuer:      h.config.Issuer(),
		AccountName: email,
	}

	key, err := totp.Generate(opts)
	if err != nil {
		return "", "", err
	}

	return key.String(), key.Secret(), nil
}
