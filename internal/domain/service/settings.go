package service

import "context"

type SettingsService interface {
	ReceiveMfaStatus(ctx context.Context, sessionKey string) bool
	DisableMfa(ctx context.Context, sessionKey string) error
	ChangePassword(ctx context.Context, sessionKey string, passwordOld string, passwordNew string) error
}
