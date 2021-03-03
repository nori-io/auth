package secret

import (
	"context"

	service2 "github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SecretService struct {
}

func New() service2.SecretService {
	return &SecretService{}
}

func (srv *SecretService) PutSecret(
	ctx context.Context, data *service2.SecretData, session entity.Session) (
	login string, issuer string, err error) {
	if err := data.Validate(); err != nil {
		return "", "", err
	}

	var mfaSecret *entity.MfaSecret

	mfaSecret = &entity.MfaSecret{
		UserID: session.UserID,
		Secret: data.Secret,
	}

	if err := srv.mfaSecretRepository.Create(ctx, mfaSecret); err != nil {
		return "", "", err
	}

	userData, err := srv.userRepository.Get(ctx, session.UserID)
	if err != nil {
		return "", "", err
	}

	if userData.Email != "" {
		login = userData.Email
	} else {
		login = userData.PhoneCountryCode + userData.PhoneNumber
	}
	return login, srv.config.Issuer, nil
}
