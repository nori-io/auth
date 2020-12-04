package http

import (
	"github.com/nori-io/authentication/internal/domain/service"
	"github.com/nori-io/authentication/internal/handler/http/authentication"
	"github.com/nori-io/http/pkg"
)

type handler struct {
	r         pkg.Http
	auth      service.AuthenticationService
	urlPrefix string
}

func New(h handler) {
	authHandler := authentication.New(h.auth)

	// todo: add middleware

	h.r.Get("/signup", authHandler.SignUp)
	h.r.Get("/signin", authHandler.SigIn)
	h.r.Get("/signout", authHandler.SignOut)
}
