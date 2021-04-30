package mfa_totp

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_totp"
)

type mfaTotpHelper struct {
	config config.Config
}

type Params struct {
	Config config.Config
}

func New(params Params) mfa_totp.MfaTotpHelper {
	return mfaTotpHelper{config: params.Config}
}
