package helper

import "github.com/nori-plugins/authentication/internal/domain/entity"

type GothProviderHelper interface {
	Use(provider *entity.SocialProvider)
	UseAll(providers []entity.SocialProvider)
}
