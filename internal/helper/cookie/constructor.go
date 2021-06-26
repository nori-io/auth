package cookie

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper"
)

type CookieHelper struct {
	config config.Config
}

type Params struct {
	Config config.Config
}

func New(params Params) helper.CookieHelper {
	return &CookieHelper{
		config: params.Config,
	}
}
