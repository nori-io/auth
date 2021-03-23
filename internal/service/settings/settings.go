package settings

import "context"

func (s SettingsService) ReceiveMfaStatus(ctx context.Context, sessionKey string) bool {
	panic("implement me")
}

func (s SettingsService) DisableMfa(ctx context.Context, sessionKey string) error {
	panic("implement me")
}

func (s SettingsService) ChangePassword(ctx context.Context, sessionKey string, passwordOld string, passwordNew string) error {
	// расшифровать пароль с salt и hash и установить в качестве нового
	panic("implement me")
}
