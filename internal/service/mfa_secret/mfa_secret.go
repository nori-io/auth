package mfa_secret

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv *MfaSecretService) PutSecret(ctx context.Context, data service.SecretData) (
	email string, issuer string, err error) {
	if err := data.Validate(); err != nil {
		return "", "", err
	}

	var mfaSecret *entity.MfaSecret

	if err := srv.mfaSecretRepository.Create(ctx, mfaSecret); err != nil {
		return "", "", err
	}

	user, err := srv.userService.GetByID(ctx, service.GetByIdData{Id: 0})
	if err != nil {
		return "", "", err
	}

	mfaSecret = &entity.MfaSecret{
		UserID: user.ID,
		Secret: data.Secret,
	}
	if user.Email != "" {
		email = user.Email
	} else {
		email = user.PhoneCountryCode + user.PhoneNumber
	}
	return email, srv.config.Issuer(), nil
}
