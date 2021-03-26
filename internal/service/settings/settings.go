package settings

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	s "github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv SettingsService) ReceiveMfaStatus(ctx context.Context, sessionKey string) (*bool, error) {
	if err := srv.session.Get([]byte(sessionKey), s.SessionActive); err != nil {
		return nil, err
	}

	session, err := srv.sessionRepository.FindBySessionKey(ctx, sessionKey)
	if err != nil {
		return nil, err
	}

	user, err := srv.userRepository.FindById(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	mfaEnabled := user.MfaType.String() != "None"

	return &mfaEnabled, err
}

func (srv SettingsService) DisableMfa(ctx context.Context, sessionKey string) error {
	if err := srv.session.Get([]byte(sessionKey), s.SessionActive); err != nil {
		return err
	}

	session, err := srv.sessionRepository.FindBySessionKey(ctx, sessionKey)
	if err != nil {
		return err
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
	if err := srv.session.Get([]byte(sessionKey), s.SessionActive); err != nil {
		return err
	}

	session, err := srv.sessionRepository.FindBySessionKey(ctx, sessionKey)
	if err != nil {
		return err
	}
	user, err := srv.userRepository.FindById(ctx, session.UserID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordOld)); err != nil {
		return err
	}

	if err := srv.userRepository.Update(ctx, &entity.User{
		ID:       user.ID,
		Password: passwordNew,
	}); err != nil {
		return err
	}

	return nil
}
