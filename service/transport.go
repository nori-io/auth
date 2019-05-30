package service

import (
	"github.com/nori-io/nori-common/endpoint"
	"github.com/nori-io/nori-common/logger"
	"github.com/nori-io/nori-interfaces/interfaces"
	"github.com/nori-io/nori-interfaces/transport/http"
)

type PluginParameters struct {
	UserTypeParameter                  []interface{}
	UserTypeDefaultParameter           string
	UserRegistrationByPhoneNumber      bool
	UserRegistrationByEmailAddress     bool
	UserMfaTypeParameter               string
	ActivationCode                     bool
	ActivationTimeForActivationMinutes uint
}

func Transport(
	auth interfaces.Auth,
	transport interfaces.HTTPTransport,
	session interfaces.Session,
	router interfaces.Http,
	srv Service,
	logger logger.Writer,
	parameters PluginParameters,

) {

	authenticated := func(e endpoint.Endpoint) endpoint.Endpoint {
		return auth.Authenticated()(session.Verify()(e))
	}

	signupHandler := http.NewServer(
		MakeSignUpEndpoint(srv, parameters),
		DecodeSignUpRequest(PluginParameters{
			UserTypeParameter:                  parameters.UserTypeParameter,
			UserTypeDefaultParameter:           parameters.UserTypeDefaultParameter,
			UserRegistrationByPhoneNumber:      parameters.UserRegistrationByPhoneNumber,
			UserRegistrationByEmailAddress:     parameters.UserRegistrationByEmailAddress,
			UserMfaTypeParameter:               parameters.UserMfaTypeParameter,
			ActivationTimeForActivationMinutes: parameters.ActivationTimeForActivationMinutes,
			ActivationCode:                     parameters.ActivationCode,
		}),
		http.EncodeJSONResponse,
	)

	http.ServerErrorLogger(logger)(signupHandler)

	signinHandler := http.NewServer(
		MakeSignInEndpoint(srv, parameters),
		DecodeSignInRequest,
		http.EncodeJSONResponse,
	)

	http.ServerErrorLogger(logger)(signinHandler)

	opts := []http.ServerOption{
		http.ServerBefore(transport.ToContext()),
	}

	signoutHandler := http.NewServer(
		authenticated(MakeSignOutEndpoint(srv)),
		DecodeSignOutRequest,
		http.EncodeJSONResponse,
		opts...,
	)

	http.ServerErrorLogger(logger)(signoutHandler)

	recoveryCodesHandler := http.NewServer(
		MakeRecoveryCodesEndpoint(srv),
		DecodeRecoveryCodes(),
		http.EncodeJSONResponse,
	)
	http.ServerErrorLogger(logger)(recoveryCodesHandler)

	router.Handle("/auth/signup", signupHandler).Methods("POST")
	router.Handle("/auth/signin", signinHandler).Methods("POST")
	router.Handle("/auth/signout", signoutHandler).Methods("GET")
	router.Handle("/auth/settings/two_factor_authentication/recovery_codes", recoveryCodesHandler).Methods("GET")

	//	/auth/verify/(uuid)
	//	/auth/delete
	//	/auth/profile
	// /auth/forgotpassword
	// /auth/resetpassword/(uuid)

}
