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

func New(h Handler) {
	authHandler := authentication.New(h.Auth)

	// todo: add middleware

	h.R.Get("/signup", authHandler.SignUp)
	h.R.Get("/signin", authHandler.SigIn)
	h.R.Get("/signout", authHandler.SignOut)
	h.R.Get("/mfa/recovery_codes", authHandler.GetMfaRecoveryCodes)
	// h.R.Put("/mfa/recovery_codes", authHandler.MfaRecoveryCodes)
}
