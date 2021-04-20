package social_provider

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"
	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SocialProviderHandler struct {
	socialProviderService service.SocialProvider
	logger                logger.FieldLogger
	cookieHelper          cookie.CookieHelper
	errorHelper           error2.ErrorHelper
}

type Params struct {
	SocialProviderService service.SocialProvider
	Logger                logger.FieldLogger
	CookieHelper          cookie.CookieHelper
	ErrorHelper           error2.ErrorHelper
}

func New(params Params) *SocialProviderHandler {
	return &SocialProviderHandler{
		socialProviderService: params.SocialProviderService,
		logger:                params.Logger,
		cookieHelper:          params.CookieHelper,
		errorHelper:           params.ErrorHelper,
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



