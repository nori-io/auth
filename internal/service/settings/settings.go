package settings

import "context"

func (s SettingsService) ReceiveMfaStatus(ctx context.Context, sessionKey string) bool {
	panic("implement me")
}

func (s SettingsService) DisableMfa(ctx context.Context, sessionKey string) error {
	panic("implement me")
}

func (s SettingsService) ChangePassword(ctx context.Context, sessionKey string, passwordOld string, passwordNew string) error {
	panic("implement me")
}
