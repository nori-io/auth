package service

import (
	"context"
	http3 "net/http"
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/openidConnect"
	"github.com/markbates/goth/providers/vk"
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
	Oath2ProvidersVKClientKey string
	Oath2ProvidersVKClientSecret string
	Oath2ProvidersVKRedirectUrl string
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


	if (len(parameters.Oath2ProvidersVKClientKey)==0)&&(len(parameters.Oath2ProvidersVKClientSecret)==0){
		goth.UseProviders(
			vk.New(os.Getenv("VK_KEY"), os.Getenv("VK_SECRET"), "http://localhost:3000/auth/vk/callback"))
		}

	openidConnect, _ := openidConnect.New(os.Getenv("OPENID_CONNECT_KEY"), os.Getenv("OPENID_CONNECT_SECRET"), "http://localhost:3000/auth/openid-connect/callback", os.Getenv("OPENID_CONNECT_DISCOVERY_URL"))
	if openidConnect != nil {
		goth.UseProviders(openidConnect)
	}



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


	signInSocialHandler:=http.NewServer(
		MakeSignInSocialEndpoint(srv, parameters),
		DecodeSignInSocial(PluginParameters{
			Oath2ProvidersVKClientSecret:parameters.Oath2ProvidersVKClientSecret,
			Oath2ProvidersVKClientKey:parameters.Oath2ProvidersVKClientKey,
			Oath2ProvidersVKRedirectUrl:parameters.Oath2ProvidersVKRedirectUrl,
		}),
		http.EncodeJSONResponse,
		)
	//http.ServerErrorHandler(logger)(signInSocialHandler)

	signOutSocialHandler:=http.NewServer(
		MakeSignOutSocial(srv),
		DecodeSignOutSocial(),
		http.EncodeJSONResponse,)
	//http.ServerErrorHandler(logger)(signOutSocialHandler)




	router.Handle("/auth/signup", signupHandler).Methods("POST")
	router.Handle("/auth/signin", signinHandler).Methods("POST")
	router.Handle("/auth/signout", signoutHandler).Methods("GET")
	router.Handle("/auth/settings/two_factor_authentication/recovery_codes", recoveryCodesHandler).Methods("GET")

	router.HandleFunc("auth/{provider}/signin", func(res http3.ResponseWriter, req *http3.Request) {
		srv.SignOutSocial(res, req)
	}).Methods("GET")


	router.HandleFunc("auth/{provider}/signout", func(res http3.ResponseWriter, req *http3.Request) {
		srv.SignInSocial(context.Background(), req, parameters)
	}).Methods("POST")

	/*/auth/{provider}
	/auth/{provider}/callback
	/logout/{provider}*/

	//	/auth/verify/(uuid)
	//	/auth/delete
	//	/auth/profile
	// /auth/forgotpassword
	// /auth/resetpassword/(uuid)

}
