package service

import (
	"github.com/nori-io/nori-common/interfaces"
	"github.com/sirupsen/logrus"
	"github.com/nori-io/nori-common/transport/http"
)

func Transport(
	transport interfaces.HTTPTransport,
	router interfaces.Http,
	srv Service,
	logger *logrus.Logger,
) {


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

		router.Handle("/v1/auth/signup", signupHandler).Methods("POST")
		router.Handle("/v1/auth/login", loginHandler).Methods("POST")
}
