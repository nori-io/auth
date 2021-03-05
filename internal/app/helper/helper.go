package helper

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/helper"
	mfaRecoveryCodeHelper "github.com/nori-plugins/authentication/internal/helper/mfa_recovery_code"
)

var HelperSet = wire.NewSet(
	wire.Struct(new(mfaRecoveryCodeHelper.Params), "Config"),
	mfaRecoveryCodeHelper.New,
	helper.New,
)
