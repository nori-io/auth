package goth_provider

import (
	"github.com/nori-plugins/authentication/internal/domain/helper"
)

type GothProviderHelper struct {
}

func New() helper.GothProviderHelper {
	return &GothProviderHelper{}
}
