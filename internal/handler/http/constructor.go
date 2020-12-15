package http_handler

import (
	"github.com/nori-io/authentication/internal/domain/service"
	"github.com/nori-io/authentication/internal/handler/http/authentication"
	httpPlugin "github.com/nori-io/interfaces/nori/http"
)

type Handler struct {
	R         httpPlugin.Http
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
	h.R.Get("/signup", authHandler.SignUp)
	h.R.Get("/signin", authHandler.SigIn)
	h.R.Get("/signout", authHandler.SignOut)
}
