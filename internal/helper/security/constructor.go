package security

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
)

type securityHelper struct {
	config config.Config
}

type Params struct {
	Config config.Config
}

func New(params Params) security.SecurityHelper {
	return &securityHelper{config: params.Config}
}
