package mfa_secret

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv *MfaSecretService) PutSecret(ctx context.Context, data *service.SecretData, session entity.Session) (
	login string, issuer string, err error) {
	if err := data.Validate(); err != nil {
		return "", "", err
	}

	var mfaSecret *entity.MfaSecret

	mfaSecret = &entity.MfaSecret{
		UserID: session.UserID,
		Secret: data.Secret,
	}

	if err := srv.MfaSecretRepository.Create(ctx, mfaSecret); err != nil {
		return "", "", err
	}

	userData, err := srv.UserRepository.FindById(ctx, session.UserID)
	if err != nil {
		return "", "", err
	}

	if userData.Email != "" {
		login = userData.Email
	} else {
		login = userData.PhoneCountryCode + userData.PhoneNumber
	}
	return login, srv.Config.Issuer(), nil
}
