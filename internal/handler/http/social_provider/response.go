package social_provider

import "github.com/nori-plugins/authentication/internal/domain/entity"

type SocialProviderResponse struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}

func convertAll(entities []entity.SocialProvider) []SocialProviderResponse {
	socialProviders := make([]SocialProviderResponse, 0)
	for _, v := range entities {
		socialProviders = append(socialProviders, convert(v))
	}
	return socialProviders
}

func convert(e entity.SocialProvider) SocialProviderResponse {
	return SocialProviderResponse{
		Name: e.Name,
		Logo: e.Logo,
	}
}
