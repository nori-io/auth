package helper

import (
	"github.com/google/wire"
	mfaRecoveryCodeHelper "github.com/nori-plugins/authentication/internal/helper/mfa_recovery_codes"
)

var HelperSet = wire.NewSet(
	mfaRecoveryCodeHelper.New,
)
