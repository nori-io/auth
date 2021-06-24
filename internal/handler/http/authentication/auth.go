package authentication

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	http2 "github.com/nori-io/interfaces/nori/http/v2"

	"github.com/markbates/goth/gothic"
	"github.com/nori-plugins/authentication/internal/domain/errors"
	"github.com/nori-plugins/authentication/pkg/enum/social_provider_status"

	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"

	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
	"github.com/nori-plugins/authentication/pkg/enum/session_status"

	"github.com/nori-plugins/authentication/internal/config"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-io/common/v5/pkg/domain/logger"

	"github.com/nori-plugins/authentication/internal/domain/entity"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

type AuthenticationHandler struct {
	R                     http2.Router
	authenticationService service.AuthenticationService
	sessionService        service.SessionService
	socialProviderService service.SocialProvider
	cookieHelper          cookie.CookieHelper
	errorHelper           error2.ErrorHelper
	config                config.Config
	logger                logger.FieldLogger
}

type Params struct {
	R                     http2.Router
	AuthenticationService service.AuthenticationService
	SessionService        service.SessionService
	SocialProviderService service.SocialProvider
	CookieHelper          cookie.CookieHelper
	ErrorHelper           error2.ErrorHelper
	Config                config.Config
	Logger                logger.FieldLogger
}

func New(params Params) *AuthenticationHandler {
	return &AuthenticationHandler{
		R:                     params.R,
		authenticationService: params.AuthenticationService,
		sessionService:        params.SessionService,
		socialProviderService: params.SocialProviderService,
		cookieHelper:          params.CookieHelper,
		errorHelper:           params.ErrorHelper,
		config:                params.Config,
		logger:                params.Logger,
	}
}

func (h *AuthenticationHandler) Session(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	sess, user, err := h.authenticationService.GetSessionData(r.Context(), service.GetSessionData{SessionKey: sessionId})
	if err != nil {
		h.logger.Error("%s", err)
		h.errorHelper.Error(w, err)
	}

	response.JSON(w, r, SessionResponse{
		Success:  true,
		Message:  "session exists",
		Email:    user.Email,
		Phone:    user.PhoneCountryCode + user.PhoneNumber,
		OpenedAt: sess.OpenedAt,
	})
}

func (h *AuthenticationHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	data, err := newSignUpData(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = h.authenticationService.SignUp(r.Context(), data)
	if err != nil {
		h.logger.Error("%s", err)
		h.errorHelper.Error(w, err)
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AuthenticationHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	data, err := newLogInData(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sess, mfaType, err := h.authenticationService.LogIn(r.Context(), data)
	if err != nil {
		h.logger.Error("%s", err)
		h.errorHelper.Error(w, err)
	}

	h.cookieHelper.SetSession(w, sess)

	response.JSON(w, r, LogInResponse{
		Success: true,
		Message: "User sign in",
		MfaType: *mfaType,
	})
}

func (h *AuthenticationHandler) LogInMfa(w http.ResponseWriter, r *http.Request) {
	_, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	data, err := newLogInMfaData(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sess, err := h.authenticationService.LogInMfa(r.Context(), data)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	h.cookieHelper.SetSession(w, sess)

	response.JSON(w, r, LogInMfaResponse{
		Success: true,
		Message: "User sign in by mfa",
	})
}

func (h *AuthenticationHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	data := &entity.Session{
		ID:         0,
		UserID:     0,
		SessionKey: nil,
		Status:     0,
		OpenedAt:   time.Time{},
		ClosedAt:   time.Time{},
		UpdatedAt:  time.Time{},
	}
	//@todo
	if data.Status != session_status.Active {
	}

	if err := h.authenticationService.LogOut(r.Context(), service.LogOutData{SessionKey: sessionId}); err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	h.cookieHelper.UnsetSession(w)

	// todo: redirect
	http.Redirect(w, r, h.config.UrlLogoutRedirect(), 0)
}

func (h *AuthenticationHandler) HandleSocialProvider(w http.ResponseWriter, r *http.Request) {
	name := h.R.URLParam(r, "social_provider")

	h.checkProviderName(w, r, name)

	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(w, gothUser)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (h *AuthenticationHandler) HandleSocialProviderCallBack(w http.ResponseWriter, r *http.Request) {
	name := h.R.URLParam(r, "social_provider")

	h.checkProviderName(w, r, name)

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(w, user)
}

func (h *AuthenticationHandler) HandleSocialProviderLogout(w http.ResponseWriter, r *http.Request) {
	name := h.R.URLParam(r, "social_provider")

	h.checkProviderName(w, r, name)

	gothic.Logout(w, r)
	w.Header().Set("Location", h.config.UrlLogoutRedirect())
	w.WriteHeader(http.StatusTemporaryRedirect)
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

func (h *AuthenticationHandler) checkProviderName(w http.ResponseWriter, r *http.Request, name string) {
	data := service.GetByNameData{Name: name}

	provider, err := h.socialProviderService.GetByName(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if provider.Status != social_provider_status.Enabled {
		http.Error(w, errors.SocialProviderNotFound.Error(), http.StatusBadRequest)
	}
}
