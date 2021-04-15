package settings

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/errors"
)

func (srv SettingsService) ReceiveMfaStatus(ctx context.Context, sessionKey string) (*bool, error) {
	session, err := srv.sessionRepository.FindBySessionKey(ctx, sessionKey)
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

func (srv SettingsService) DisableMfa(ctx context.Context, sessionKey string) error {
	session, err := srv.sessionRepository.FindBySessionKey(ctx, sessionKey)
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

func (srv SettingsService) ChangePassword(ctx context.Context, sessionKey string, passwordOld string, passwordNew string) error {
	session, err := srv.sessionRepository.FindBySessionKey(ctx, sessionKey)
	if err != nil {
		return err
	}
	if session == nil {
		return errors.SessionNotFound
	}
	user, err := srv.userService.GetByID(ctx, session.UserID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.UserNotFound
	}

	if err := srv.securityHelper.ComparePassword(passwordOld, user.Password); err != nil {
		return err
	}

	if err := srv.userService.UpdatePassword(ctx, service.UserUpdatePasswordData{
		UserID:   session.UserID,
		Password: passwordNew,
	}); err != nil {
		return err
	}

	return nil
}
