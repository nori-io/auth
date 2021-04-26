package settings

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/errors"
)

func (srv SettingsService) ReceiveMfaStatus(ctx context.Context, data service.ReceiveMfaStatusData) (*bool, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	session, err := srv.sessionRepository.FindBySessionKey(ctx, data.SessionKey)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, errors.SessionNotFound
	}

	user, err := srv.userService.GetByID(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.UserNotFound
	}

	mfaEnabled := user.MfaType.String() != "None"

	return &mfaEnabled, err
}

func (srv SettingsService) DisableMfa(ctx context.Context, data service.DisableMfaData) error {
	if err := data.Validate(); err != nil {
		return err
	}

	session, err := srv.sessionRepository.FindBySessionKey(ctx, data.SessionKey)
	if err != nil {
		return err
	}

	if session == nil {
		return errors.SessionNotFound
	}

	if err := srv.userService.UpdateMfaStatus(ctx, service.UserUpdateMfaStatusData{
		UserID:  session.UserID,
		MfaType: 0,
	}); err != nil {
		return err
	}

	return nil
}

func (srv SettingsService) ChangePassword(ctx context.Context, data service.ChangePasswordData) error {
	if err := data.Validate(); err != nil {
		return err
	}

	session, err := srv.sessionRepository.FindBySessionKey(ctx, data.SessionKey)
	if err != nil {
		return err
	}
	if session == nil {
		return errors.SessionNotFound
	}
	user, err := srv.userService.GetByID(ctx, service.GetByIdData{Id: session.UserID})
	if err != nil {
		return err
	}
	if user == nil {
		return errors.UserNotFound
	}

	if err := srv.securityHelper.ComparePassword(data.PasswordOld, user.Password); err != nil {
		return err
	}

	if err := srv.userService.UpdatePassword(ctx, service.UserUpdatePasswordData{
		UserID:   session.UserID,
		Password: data.PasswordNew,
	}); err != nil {
		return err
	}

	return nil
}
