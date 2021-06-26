package mfa_totp

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper"
)

type mfaTotpHelper struct {
	config config.Config
}

type Params struct {
	Config config.Config
}

func New(params Params) helper.MfaTotpHelper {
	return mfaTotpHelper{config: params.Config}
}
