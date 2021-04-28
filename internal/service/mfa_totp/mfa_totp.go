package mfa_totp

import (
	"context"
	"time"

	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"

	"github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv *MfaTotpService) GetSecret(ctx context.Context, data service.SecretData) (
	string, error) {
	if err := data.Validate(); err != nil {
		return "", err
	}

	session, err := srv.sessionService.GetBySessionKey(ctx, service.GetBySessionKeyData{SessionKey: data.SessionKey})
	if err != nil {
		return "", err
	}

	if session == nil {
		return "", errors2.SessionNotFound
	}

	user, err := srv.userService.GetByID(ctx, service.GetByIdData{Id: session.UserID})
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors2.UserNotFound
	}

	secret, err := srv.totpHelper.Generate(user.Email)
	if err != nil {
		return "", err
	}

	if err := srv.mfaSecretRepository.Delete(ctx, user.ID); err != nil {
		return "", nil
	}

	if err := srv.mfaSecretRepository.Create(ctx, &entity.MfaSecret{
		UserID:    user.ID,
		Secret:    secret,
		CreatedAt: time.Now(),
	}); err != nil {
		return "", err
	}

	return secret, nil
}
