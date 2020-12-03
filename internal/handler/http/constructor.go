package http

import (
	"github.com/nori-io/authentication/internal/domain/service"
	"github.com/nori-io/authentication/internal/handler/http/authentication"
	"github.com/nori-io/http/pkg"
)

func New(r pkg.Http, auth service.AuthenticationService) {
	authHandler := authentication.New(auth)

	// todo: add middleware

	r.Get("/signup", authHandler.SignUp)
	r.Get("/signin", authHandler.SigIn)
	r.Get("/signout", authHandler.SignOut)
}
