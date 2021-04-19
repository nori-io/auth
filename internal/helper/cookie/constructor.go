package cookie

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
)

type CookieHelper struct {
	config config.Config
}

type Params struct {
	Config config.Config
}

func New(params Params) cookie.CookieHelper {
	return &CookieHelper{
		config: params.Config,
	}
}
