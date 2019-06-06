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

	type ProviderIndex struct {
		Providers    []string
		ProvidersMap map[string]string
	}

	if (len(parameters.Oath2ProvidersVKClientKey) == 0) && (len(parameters.Oath2ProvidersVKClientSecret) == 0) {
		goth.UseProviders(
			vk.New(parameters.Oath2ProvidersVKClientKey, parameters.Oath2ProvidersVKClientSecret, parameters.Oath2ProvidersVKRedirectUrl))
	}

	m := make(map[string]string)

	/*	openidConnect, _ := openidConnect.New(os.Getenv("OPENID_CONNECT_KEY"), os.Getenv("OPENID_CONNECT_SECRET"), "http://localhost:3000/auth/openid-connect/callback", os.Getenv("OPENID_CONNECT_DISCOVERY_URL"))
		if openidConnect != nil {
			goth.UseProviders(openidConnect)
		}*/

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	m["vk"] = "VK"
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

	/*signInSocialHandler := http.NewServer(
		MakeSignInSocialEndpoint(srv, parameters),
		DecodeSignInSocial(PluginParameters{
			Oath2ProvidersVKClientSecret: parameters.Oath2ProvidersVKClientSecret,
			Oath2ProvidersVKClientKey:    parameters.Oath2ProvidersVKClientKey,
			Oath2ProvidersVKRedirectUrl:  parameters.Oath2ProvidersVKRedirectUrl,
		}),
		http.EncodeJSONResponse,
	)
	//http.ServerErrorHandler(logger)(signInSocialHandler)

	signOutSocialHandler := http.NewServer(
		MakeSignOutSocial(srv),
		DecodeSignOutSocial(),
		http.EncodeJSONResponse)
	//http.ServerErrorHandler(logger)(signOutSocialHandler)*/

	router.Handle("/auth/signup", signupHandler).Methods("POST")
	router.Handle("/auth/signin", signinHandler).Methods("POST")
	router.Handle("/auth/signout", signoutHandler).Methods("GET")
	router.Handle("/auth/settings/two_factor_authentication/recovery_codes", recoveryCodesHandler).Methods("GET")

	router.HandleFunc("/auth/{provider}", func(res httpNet.ResponseWriter, req *httpNet.Request) {
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			t, _ := template.New("foo").Parse(userTemplate)
			t.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	}).Methods("GET")

	router.HandleFunc("/auth/{provider}/callback", func(res httpNet.ResponseWriter, req *httpNet.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
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


var indexTemplate = `{{range $key,$value:=.Providers}}
    <p><a href="/auth/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}`

var userTemplate = `
<p><a href="/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`