package mfa_totp

import (
	"context"
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/users_action"

	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"

	"github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv *MfaTotpService) GetUrl(ctx context.Context, data service.MfaGetUrlData) (
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

	url, secret, err := srv.mfaTotpHelper.Generate(user.Email)
	if err != nil {
		return "", err
	}

	if err := srv.transactor.Transact(ctx, func(tx context.Context) error {
		if err := srv.mfaTotpRepository.Delete(ctx, user.ID); err != nil {
			return nil
		}

		if err := srv.mfaTotpRepository.Create(ctx, &entity.MfaTotp{
			UserID:    user.ID,
			Secret:    secret,
			CreatedAt: time.Now(),
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return "", err
	}

	return url, nil
}

func (srv *MfaTotpService) Validate(ctx context.Context, data service.MfaTotpValidateData) (bool, error) {
	if err := data.Validate(); err != nil {
		return false, err
	}

	session, err := srv.sessionService.GetBySessionKey(ctx, service.GetBySessionKeyData{SessionKey: data.SessionKey})
	if err != nil {
		return false, err
	}

	if session == nil {
		return false, errors2.SessionNotFound
	}

	mfaTotp, err := srv.mfaTotpRepository.FindByUserId(ctx, data.UserID)
	if err != nil {
		return false, err
	}
	if mfaTotp == nil {
		return false, errors2.MfaTotpSecretNotFound
	}

	isTokenValid := srv.mfaTotpHelper.Validate(data.PassCode, mfaTotp.Secret)

	if isTokenValid {
		if err := srv.userLogService.Create(ctx, service.UserLogCreateData{
			UserID:    mfaTotp.UserID,
			Action:    users_action.MfaOtpEnabled,
			SessionID: session.ID,
			Meta:      "",
			CreatedAt: time.Now(),
		}); err != nil {
			return false, err
		}
	}

	return isTokenValid, nil
}
