package mfa_totp

import "github.com/nori-plugins/authentication/internal/config"

type totpHelper struct {
	config config.Config
}

type Params struct {
	config config.Config
}

func New(params Params) totpHelper {
	return totpHelper{config: params.config}
}
