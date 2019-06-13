package service

import (
	"fmt"
	"html/template"
	httpNet "net/http"
	"sort"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
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
	Oath2ProvidersVKClientKey          string
	Oath2ProvidersVKClientSecret       string
	Oath2ProvidersVKRedirectUrl        string
	Oath2SessionSecret                 string
}

var Oath2SessionSecret string

func Transport(
	auth interfaces.Auth,
	transport interfaces.HTTPTransport,
	session interfaces.Session,
	router interfaces.Http,
	srv Service,
	logger logger.Writer,
	parameters PluginParameters,

) {

	if (len(parameters.Oath2ProvidersVKClientKey) > 0) && (len(parameters.Oath2ProvidersVKClientSecret) > 0) {
		goth.UseProviders(
			vk.New(parameters.Oath2ProvidersVKClientKey, parameters.Oath2ProvidersVKClientSecret, parameters.Oath2ProvidersVKRedirectUrl))
	}

	Oath2SessionSecret = parameters.Oath2SessionSecret
	Init(Oath2SessionSecret)

	type ProviderIndex struct {
		Providers    []string
		ProvidersMap map[string]string
	}
	m := make(map[string]string)
	m["vk"] = "VK"
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}

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

	/*	signInSocialHandler := http.NewServer(
			MakeSignInSocialEndpoint(srv, parameters),
			DecodeSignInSocial,
			http.EncodeJSONResponse,
		)
		http.ServerErrorHandler(logger)(signInSocialHandler)

		signOutSocialHandler := http.NewServer(
			MakeSignOutSocial(srv),
			DecodeSocialSignOut,
			http.EncodeJSONResponse)
		http.ServerErrorHandler(logger)(signOutSocialHandler)*/

	router.Handle("/auth/signup", signupHandler).Methods("POST")
	router.Handle("/auth/signin", signinHandler).Methods("POST")
	router.Handle("/auth/signout", signoutHandler).Methods("GET")
	router.Handle("/auth/settings/two_factor_authentication/recovery_codes", recoveryCodesHandler).Methods("GET")

	router.HandleFunc("/auth/{provider}", func(res httpNet.ResponseWriter, req *httpNet.Request) {
		fmt.Println("/auth/{provider}", "here")
		srv.SignInSocial(res, *req)
	}).Methods("GET")

	router.HandleFunc("/auth/{provider}/callback", func(res httpNet.ResponseWriter, req *httpNet.Request) {

		user, err := CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, user)
	}).Methods("GET")

	router.HandleFunc("/logout/{provider}", func(res httpNet.ResponseWriter, req *httpNet.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(httpNet.StatusTemporaryRedirect)
	}).Methods("GET")

	router.HandleFunc("/", func(res httpNet.ResponseWriter, req *httpNet.Request) {
		t, _ := template.New("foo").Parse(indexTemplate)
		t.Execute(res, providerIndex)
	}).Methods("GET")
	logger.Println("listening on localhost:8080")
	logger.Error(httpNet.ListenAndServe(":8080 error", signinHandler))

}
