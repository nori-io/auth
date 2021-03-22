package settings

import "context"

func (s SettingsService) ReceiveMfaStatus(ctx context.Context, sessionKey string) {
	panic("implement me")
}

func (s SettingsService) DisableMfa(ctx context.Context, sessionKey string) {
	panic("implement me")
}

func (s SettingsService) ChangePassword(ctx context.Context, sessionKey string, passwordOld string, passwordNew string) {
	panic("implement me")
}
