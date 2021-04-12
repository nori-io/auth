package settings

import (
	"context"

	errors2 "github.com/nori-plugins/authentication/pkg/errors"

	"github.com/nori-plugins/authentication/internal/domain/errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv SettingsService) ReceiveMfaStatus(ctx context.Context, sessionKey string) (*bool, error) {
	session, err := srv.sessionRepository.FindBySessionKey(ctx, sessionKey)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, errors.SessionNotFound
	}

	user, err := srv.userRepository.FindById(ctx, session.UserID)
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

	if err := srv.userRepository.Update(ctx, &entity.User{
		ID:      session.UserID,
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
	user, err := srv.userRepository.FindById(ctx, session.UserID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.UserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordOld)); err != nil {
		return errors2.NewInternal(err)
	}

	if err := srv.userRepository.Update(ctx, &entity.User{
		ID:       user.ID,
		Password: passwordNew,
	}); err != nil {
		return err
	}

	return nil
}
