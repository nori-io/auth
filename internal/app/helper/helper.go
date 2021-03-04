package helper

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/helper"
	mfaRecoveryCodeHelper "github.com/nori-plugins/authentication/internal/helper/mfa_recovery_code"
)

var HelperSet = wire.NewSet(
	mfaRecoveryCodeHelper.New,
	helper.New,
)
