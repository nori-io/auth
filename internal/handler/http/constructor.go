package http_handler

import (
	"github.com/nori-io/authentication/internal/domain/service"
	"github.com/nori-io/authentication/internal/handler/http_handler/authentication"
	httpPlugin "github.com/nori-io/http"
)

type Handler struct {
	r         httpPlugin1.
	auth      service.AuthenticationService
	urlPrefix string
}

func New(h Handler) *Handler {
	handler := Handler{
		r:         h.r,
		auth:      h.auth,
		urlPrefix: h.urlPrefix,
	}
	// todo: add middleware
	Start(h)
	return &handler
}

func Start(h Handler) {
	authHandler := authentication.New(h.auth)
	h.r.Get("/signup", authHandler.SignUp)
	h.r.Get("/signin", authHandler.SigIn)
	h.r.Get("/signout", authHandler.SignOut)
}
