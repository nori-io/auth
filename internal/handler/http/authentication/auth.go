package authentication

import (
	"net/http"
	"time"

	"github.com/nori-plugins/authentication/internal/config"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-io/common/v4/pkg/domain/logger"

	"github.com/nori-plugins/authentication/internal/domain/entity"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

type AuthenticationHandler struct {
	authenticationService service.AuthenticationService
	logger                logger.FieldLogger
	config                config.Config
}

type Params struct {
	AuthenticationService service.AuthenticationService
	Logger                logger.FieldLogger
	Config                config.Config
}

func New(params Params) *AuthenticationHandler {
	return &AuthenticationHandler{
		authenticationService: params.AuthenticationService,
		logger:                params.Logger,
		config:                params.Config,
	}
}

func (h *AuthenticationHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	data, err := newSignUpData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = h.authenticationService.SignUp(r.Context(), data)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AuthenticationHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	data, err := newSignInData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sess, mfaType, err := h.authenticationService.SignIn(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	c := http.Cookie{
		Name:   "ssid",
		Value:  string(sess.SessionKey),
		Path:   h.config.CookiesPath(),
		Domain: h.config.CookiesDomain(),
		//@todo Expires
		Expires:    time.Unix(h.config.CookiesExpires(), 0),
		RawExpires: h.config.CookiesRawExpires(),
		MaxAge:     h.config.CookiesMaxAge(),
		Secure:     h.config.CookiesSecure(),
		HttpOnly:   h.config.CookiesHttpOnly(),
		SameSite:   http.SameSite(h.config.CookiesSameSite()),
		Raw:        h.config.CookiesRaw(),
		Unparsed:   h.config.CookiesUnparsed(),
	}

	http.SetCookie(w, &c)
	w.WriteHeader(http.StatusOK)

	response.JSON(w, r, SignInResponse{
		Success: true,
		Message: "User sign in",
		MfaType: *mfaType,
	})
}

func (h *AuthenticationHandler) SignInMfa(w http.ResponseWriter, r *http.Request) {
	data, err := newSignInMfaData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sess, err := h.authenticationService.SignInMfa(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	c := http.Cookie{
		Name:   "ssid",
		Value:  string(sess.SessionKey),
		Path:   h.config.CookiesPath(),
		Domain: h.config.CookiesDomain(),
		//@todo Expires
		Expires:    time.Unix(h.config.CookiesExpires(), 0),
		RawExpires: h.config.CookiesRawExpires(),
		MaxAge:     h.config.CookiesMaxAge(),
		Secure:     h.config.CookiesSecure(),
		HttpOnly:   h.config.CookiesHttpOnly(),
		SameSite:   http.SameSite(h.config.CookiesSameSite()),
		Raw:        h.config.CookiesRaw(),
		Unparsed:   h.config.CookiesUnparsed(),
	}
	http.SetCookie(w, &c)

	w.WriteHeader(http.StatusOK)

	response.JSON(w, r, SignInMfaResponse{
		Success: true,
		Message: "User sign in by mfa",
	})
}

func (h *AuthenticationHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	// todo: extract session ID from context
	sessionIdContext := r.Context().Value("session_id")

	sessionId, _ := sessionIdContext.([]byte)

	if err := h.authenticationService.SignOut(r.Context(), &entity.Session{SessionKey: sessionId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// todo: redirect

	http.Redirect(w, r, "/", 0)
}
