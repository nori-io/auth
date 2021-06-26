package social_provider

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/helper"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-io/common/v5/pkg/domain/logger"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SocialProviderHandler struct {
	socialProviderService service.SocialProvider
	cookieHelper          helper.CookieHelper
	errorHelper           helper.ErrorHelper
	logger                logger.FieldLogger
}

type Params struct {
	SocialProviderService service.SocialProvider
	CookieHelper          helper.CookieHelper
	ErrorHelper           helper.ErrorHelper
	Logger                logger.FieldLogger
}

func New(params Params) *SocialProviderHandler {
	return &SocialProviderHandler{
		socialProviderService: params.SocialProviderService,
		cookieHelper:          params.CookieHelper,
		errorHelper:           params.ErrorHelper,
		logger:                params.Logger,
	}
}

func (h *SocialProviderHandler) GetSocialProviders(w http.ResponseWriter, r *http.Request) {
	_, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	providers, err := h.socialProviderService.GetAllActive(r.Context())
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, convertAll(providers))
}
