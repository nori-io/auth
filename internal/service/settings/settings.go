package settings

import (
	"context"

	s "github.com/nori-io/interfaces/nori/session"
)

func (srv SettingsService) ReceiveMfaStatus(ctx context.Context, sessionKey string) (*bool, error) {
	err := srv.session.Get([]byte(sessionKey), s.SessionActive)
	if err != nil {
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
	panic("implement me")
}

func (srv SettingsService) ChangePassword(ctx context.Context, sessionKey string, passwordOld string, passwordNew string) error {
	// расшифровать пароль с salt и hash и установить в качестве нового
	panic("implement me")
}
