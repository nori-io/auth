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
	userType []interface{},
	userTypeDefault string,
) {

	authenticated := func(e endpoint.Endpoint) endpoint.Endpoint {
		return auth.Authenticated()(session.Verify()(e))
	}

	signupHandler := http.NewServer(
		MakeSignUpEndpoint(srv),
		DecodeSignUpRequest(userType, userTypeDefault),
		http.EncodeJSONResponse,
		logger,
	)
	signinHandler := http.NewServer(
		MakeSignInEndpoint(srv),
		DecodeLogInRequest,
		http.EncodeJSONResponse,
		logger,
	)

	opts := []http.ServerOption{
		http.ServerBefore(transport.ToContext()),
	}

	signoutHandler := http.NewServer(
		authenticated(MakeSignOutEndpoint(srv)),
		DecodeLogOutRequest,
		http.EncodeJSONResponse,
		logger,
		opts...,
	)

	router.Handle("/auth/signup", signupHandler).Methods("POST")
	router.Handle("/auth/signin", signinHandler).Methods("POST")
	router.Handle("/auth/signout", signoutHandler).Methods("GET")

}
