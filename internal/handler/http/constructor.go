package http

import (
	"github.com/nori-io/interfaces/nori/http"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/internal/handler/http/authentication"
)

type Handler struct {
	R         http.Http
	Auth      service.AuthenticationService
	UrlPrefix string
}

func New(h Handler) *Handler {
	handler := Handler{
		R:         h.R,
		Auth:      h.Auth,
		UrlPrefix: h.UrlPrefix,
	}
	// todo: add middleware
	Start(h)
	return &handler
}

func Start(h Handler) {
	authHandler := authentication.New(h.Auth)

	h.R.Get("/auth/signup", handler.SignUp)
	h.R.Get("/auth/signin", handler.SigIn)
	h.R.Get("/auth/signout", handler.SignOut)
	h.R.Get("/auth/settings/mfa/recovery_codes", handler.GetMfaRecoveryCodes)

	// mfa
	h.R.Get("/auth/settings/mfa", nil)
	h.R.Get("/auth/settings/mfa/verify?", handler.PutSecret)

	// h.R.Put("/mfa/recovery_codes", authHandler.MfaRecoveryCodes)
}
