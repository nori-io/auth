package http

import (
	"github.com/nori-io/authentication/internal/domain/service"
	"github.com/nori-io/authentication/internal/handler/http/authentication"
	"github.com/nori-io/http/pkg"
)

type Handler struct {
	R         pkg.Http
	Auth      service.AuthenticationService
	UrlPrefix string
}

func New(h Handler) {
	authHandler := authentication.New(h.Auth)

	// todo: add middleware

	h.R.Get("/signup", authHandler.SignUp)
	h.R.Get("/signin", authHandler.SigIn)
	h.R.Get("/signout", authHandler.SignOut)
}
