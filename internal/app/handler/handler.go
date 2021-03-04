package handler

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/handler/http/authentication"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_recovery_code"
)

var HandlerSet = wire.NewSet(
	authentication.New,
	mfa_recovery_code.New,
)
