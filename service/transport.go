package service

import (
	"github.com/nori-io/nori-common/endpoint"
	"github.com/nori-io/nori-common/interfaces"
	"github.com/nori-io/nori-common/transport/http"
	"github.com/sirupsen/logrus"
)

func Transport(
	auth interfaces.Auth,
	transport interfaces.HTTPTransport,
	session interfaces.Session,
	router interfaces.Http,
	srv Service,
	logger *logrus.Logger,
) {

	authenticated := func(e endpoint.Endpoint) endpoint.Endpoint {
		return auth.Authenticated()(session.Verify()(e))
	}

	signupHandler := http.NewServer(
		MakeSignUpEndpoint(srv),
		DecodeSignUpRequest,
		http.EncodeJSONResponse,
		logger,
	)

	loginHandler := http.NewServer(
		MakeLogInEndpoint(srv),
		DecodeLogInRequest,
		http.EncodeJSONResponse,
		logger,
	)

	opts := []http.ServerOption{
		http.ServerBefore(transport.ToContext()),
	}

	logoutHandler := http.NewServer(
		authenticated(MakeLogOutEndpoint(srv)),
		DecodeLogOutRequest,
		http.EncodeJSONResponse,
		logger,
		opts...,
	)

	router.Handle("/v1/auth/signup", signupHandler).Methods("POST")
	router.Handle("/v1/auth/login", loginHandler).Methods("POST")
	router.Handle("/v1/auth/logout", logoutHandler).Methods("GET")
}
