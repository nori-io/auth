package security

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper"
)

type securityHelper struct {
	config config.Config
}

type Params struct {
	Config config.Config
}

func New(params Params) helper.SecurityHelper {
	return &securityHelper{config: params.Config}
}
