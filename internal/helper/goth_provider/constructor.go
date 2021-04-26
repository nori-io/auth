package goth_provider

import "github.com/nori-plugins/authentication/internal/domain/helper/goth_provider"

type GothProviderHelper struct {
}

func New() goth_provider.GothProviderHelper {
	return &GothProviderHelper{}
}
