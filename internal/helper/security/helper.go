package security

import (
	"github.com/nori-plugins/authentication/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (h securityHelper) GenerateHash(password string) ([]byte, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), h.config.PasswordBcryptCost())
	if err != nil {
		return nil, errors.NewInternal(err)
	}
	return passwordHashed, nil
}

func (h securityHelper) ComparePassword(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return errors.NewInternal(err)
	}
	return nil
}
