package social_provider

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/markbates/goth/gothic"

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

func (h *SocialProviderHandler) HandleSocialProvider(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "social_provider")
	//@todo name is empty
	if err := h.socialProviderService.IsSocialProviderEnabled(r.Context(), name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(w, gothUser)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (h *SocialProviderHandler) HandleSocialProviderCallBack(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "social_provider")
	//@todo name is empty
	if err := h.socialProviderService.IsSocialProviderEnabled(r.Context(), name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(w, user)
}

var userTemplate = `
<p><a href="/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`
