package totp

import (
	"github.com/pquerna/otp/totp"
)

func (h totpHelper) Generate(email string) (string, error) {
	opts := totp.GenerateOpts{
		Issuer:      h.config.Issuer(),
		AccountName: email,
	}

	key, err := totp.Generate(opts)
	if err != nil {
		return "", err
	}
	return key.Secret(), nil
}
