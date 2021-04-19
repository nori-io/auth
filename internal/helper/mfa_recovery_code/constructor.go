package mfa_recovery_code

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_code"
)

type mfaRecoveryCodesHelper struct {
	config config.Config
}

type Params struct {
	Config config.Config
}

func New(params Params) mfa_recovery_code.MfaRecoveryCodesHelper {
	return &mfaRecoveryCodesHelper{config: params.Config}
}
