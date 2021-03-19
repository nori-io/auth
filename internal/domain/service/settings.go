package service

import "context"

type SettingsService interface {
	ReceiveMfaStatus(ctx context.Context, sessionKey string)
	DisableMfa(ctx context.Context, sessionKey string)
	ChangePassword(ctx context.Context, sessionKey string, passwordOld string, passwordNew string)
}
