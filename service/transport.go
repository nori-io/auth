package service

import (
	"github.com/nori-io/nori-common/interfaces"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Transport(
	//transport interfaces.HTTPTransport,
	router interfaces.Http,
	srv Service,
	logger *logrus.Logger,
) {

	/*
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
		router.Handle("/v1/auth/login", loginHandler).Methods("POST")*/


	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!!"))
	}).Methods("GET")

}
