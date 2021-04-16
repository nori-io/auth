package social_provider

import (
	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SocialProviderHandler struct {
	socialProviderService service.SocialProvider
	logger                logger.FieldLogger
}

type Params struct {
	SocialProviderService service.SocialProvider
	Logger                logger.FieldLogger
}

func New(params Params) *SocialProviderHandler {
	return &SocialProviderHandler{
		socialProviderService: params.SocialProviderService,
		logger:                params.Logger,
	}
}
