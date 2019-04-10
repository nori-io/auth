package service

import (
	"github.com/nori-io/nori-common/endpoint"
	"github.com/nori-io/nori-common/interfaces"
	"github.com/nori-io/nori-common/transport/http"

	"github.com/sirupsen/logrus"
)

type PluginParameters struct {
	UserTypeParameter                []interface{}
	UserTypeDefaultParameter         string
	UserRegistrationPhoneNumberType  bool
	UserRegistrationEmailAddressType bool
	UserMfaType string
}

func Transport(
	auth interfaces.Auth,
	transport interfaces.HTTPTransport,
	session interfaces.Session,
	router interfaces.Http,
	srv Service,
	logger *logrus.Logger,
	parameters PluginParameters,

) {

	authenticated := func(e endpoint.Endpoint) endpoint.Endpoint {
		return auth.Authenticated()(session.Verify()(e))
	}

	signupHandler := http.NewServer(
		MakeSignUpEndpoint(srv),
		DecodeSignUpRequest(PluginParameters{
			UserTypeParameter: parameters.UserTypeParameter,
			UserTypeDefaultParameter: parameters.UserTypeDefaultParameter,
			UserRegistrationPhoneNumberType: parameters.UserRegistrationPhoneNumberType,
			UserRegistrationEmailAddressType: parameters.UserRegistrationEmailAddressType,
			UserMfaType:parameters.UserMfaType,
		    }),
		http.EncodeJSONResponse,
		logger,
	)
	signinHandler := http.NewServer(
		MakeSignInEndpoint(srv),
		DecodeSignInRequest,
		http.EncodeJSONResponse,
		logger,
	)

	opts := []http.ServerOption{
		http.ServerBefore(transport.ToContext()),
	}

	signoutHandler := http.NewServer(
		authenticated(MakeSignOutEndpoint(srv)),
		DecodeSignOutRequest,
		http.EncodeJSONResponse,
		logger,
		opts...,
	)

	recoveryCodesHandler := http.NewServer(
		MakeRecoveryCodesEndpoint(srv),
		DecodeRecoveryCodes(),
		http.EncodeJSONResponse,
		logger)

	profileHandler:=http.NewServer(
		MakeProfileEndpoint(srv),
		DecodeProfile(),
		http.EncodeJSONResponse,
		logger)

	/*
		verifyHandler:=http.NewServer(
			MakeVerifyEndpoint(srv),
			DecodeVerify(),
			http.EncodeJSONResponse,
			logger)*/

	router.Handle("/auth/signup", signupHandler).Methods("POST")
	router.Handle("/auth/signin", signinHandler).Methods("POST")
	router.Handle("/auth/signout", signoutHandler).Methods("GET")
	router.Handle("auth/settings/profile", profileHandler).Methods("POST")
	router.Handle("/auth/settings/two_factor_authentication/recovery_codes", recoveryCodesHandler).Methods("POST")

	/*	router.Handle("/auth/settings/two_factor_authentication/verify",verifyHandler).Methods("POST")
	 */
	//	/auth/verify/(uuid)
	//	/auth/delete
	//	/auth/profile
	// /auth/forgotpassword
	// /auth/resetpassword/(uuid)

}
