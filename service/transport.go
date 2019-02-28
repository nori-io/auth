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

	/*	router.HandleFunc("/test", func(w http2.ResponseWriter, r *http2.Request) {
		w.Write([]byte("Hello World!!!"))
	}).Methods("GET")*/
	router.Handle("/auth/signup", signupHandler).Methods("POST")

	router.Handle("/auth/signin", signinHandler).Methods("POST")
	router.Handle("/auth/signout", signoutHandler).Methods("GET")

}
