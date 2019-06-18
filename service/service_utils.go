package service

import (
	"compress/gzip"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/nori-io/nori-interfaces/interfaces"
	"github.com/pkg/errors"
)

var keySet = false

func Init(Oath2SessionSecret string) {

	key := []byte(Oath2SessionSecret)
	keySet = len(key) != 0

}

/*
BeginAuthHandler is a convenience handler for starting the authentication process.
It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider".
BeginAuthHandler will redirect the user to the appropriate authentication end-point
for the requested provider.
See https://github.com/markbates/goth/examples/main.go to see this in action.
*/
func BeginAuthHandler(res http.ResponseWriter, req *http.Request, session interfaces.Session, sid string) {
	url, err := GetAuthURL(res, req, session, sid)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(res, err)
		return
	}

	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

// SetState sets the state string associated with the given request.
// If no state string is associated with the request, one will be generated.
// This state is sent to the provider and can be retrieved during the
// callback.
var SetState = func(req *http.Request) string {
	state := req.URL.Query().Get("state")
	if len(state) > 0 {
		return state
	}

	// If a state query param is not passed in, generate a random
	// base64-encoded nonce so that the state on the auth URL
	// is unguessable, preventing CSRF attacks, as described in
	//
	// https://auth0.com/docs/protocols/oauth2/oauth-state#keep-reading
	nonceBytes := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, nonceBytes)
	if err != nil {
		panic("gothic: source of randomness unavailable: " + err.Error())
	}
	return base64.URLEncoding.EncodeToString(nonceBytes)
}

// GetState gets the state returned by the provider during the callback.
// This is used to prevent CSRF attacks, see
// http://tools.ietf.org/html/rfc6749#section-10.12
var GetState = func(req *http.Request) string {
	return req.URL.Query().Get("state")
}

/*
GetAuthURL starts the authentication process with the requested provided.
It will return a URL that should be used to send users to.
It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider".
I would recommend using the BeginAuthHandler instead of doing all of these steps
yourself, but that's entirely up to you.
*/
func GetAuthURL(res http.ResponseWriter, req *http.Request, session interfaces.Session, sid string) (string, error) {
	if !keySet {
		fmt.Println("goth/gothic: no SESSION_SECRET environment variable is set. The default cookie store is not available and any calls will fail. Ignore this warning if you are using a different store.")
	}

	providerName, err := GetProviderName(req, session, sid )
	if err != nil {
		return "", err
	}

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return "", err
	}
	sess, err := provider.BeginAuth(SetState(req))
	if err != nil {
		return "", err
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}

	err = session.Save([]byte(providerName), "req+res",0)

	if err != nil {
		return "", err
	}

	return url, err
}

/*
CompleteUserAuth does what it says on the tin. It completes the authentication
process and fetches all of the basic information about the user from the provider.
It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider".
See https://github.com/markbates/goth/examples/main.go to see this in action.
*/
var CompleteUserAuth = func(res http.ResponseWriter, req *http.Request, session interfaces.Session, sid string) (goth.User, error) {
	defer Logout(res, req, session, sid)
	if !keySet {
		fmt.Println("goth/gothic: no SESSION_SECRET environment variable is set. The default cookie store is not available and any calls will fail. Ignore this warning if you are using a different store.")
	}

	providerName, err := GetProviderName(req, session, sid)
	if err != nil {
		return goth.User{}, err
	}
	fmt.Println("#1.1. Provide name is", providerName)
	fmt.Println("#1.10 Req is", req.Body)

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return goth.User{}, err
	}
    fmt.Println("#1.2 Provider is", provider)

	var sd sessionData

	session.Save([]byte(sid), sessionData{Provider:providerName}, 0)



    fmt.Println("#1.30 err", err)
	fmt.Println("#1.31 Sd.provider", sd.Provider)
	fmt.Println("#1.32 Sid is", sid)

	err = session.Get([]byte(sid), &sd)
	if err != nil {
		return goth.User{}, err
	}

	fmt.Println("#1.4 Value from session", sd)


	value, err := GetFromSession(providerName, req, session)
	if err != nil {
		return goth.User{}, err
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return goth.User{}, err
	}

	fmt.Println("#1.5 Session with unmarshall", sess, "  ", err)

/*	err = validateState(req, session)
	if err != nil {
		return goth.User{}, err
	}*/

	fmt.Println("#1.6 validateState", err)

	user, err := provider.FetchUser(sess)
	if err == nil {
		// user can be found with existing session data
		return user, err
	}

	// get new token and retry fetch
	_, err = sess.Authorize(provider, req.URL.Query())
	if err != nil {
		return goth.User{}, err
	}

	fmt.Println("#1.7 sess.Authorize", err)

	/*err = StoreInSession(providerName, sess.Marshal(), req, res)
	fmt.Println("StoreInSession", err)
	*/
	if err != nil {
		return goth.User{}, err
	}

	gu, err := provider.FetchUser(sess)

	fmt.Println("provider.FetchUser", gu, "     ", err)
	type sessionData struct {
		name string
	}

	return gu, err
}

/*var CompleteUserAuth = func(res http.ResponseWriter, req *http.Request) (goth.User, error) {
	defer Logout(res, req)
	if !keySet && defaultStore == Store {
		fmt.Println("goth/gothic: no SESSION_SECRET environment variable is set. The default cookie store is not available and any calls will fail. Ignore this warning if you are using a different store.")
	}

	providerName, err := GetProviderName(req)
	if err != nil {
		return goth.User{}, err
	}
	fmt.Println("Provide name is", providerName)

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return goth.User{}, err
	}

	fmt.Println("Provide is", provider)


	value, err := GetFromSession(providerName, req)
	if err != nil {
		return goth.User{}, err
	}

	fmt.Println("Value from session", value)


	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return goth.User{}, err
	}

	fmt.Println("Session with unmarshall", sess, "  ", err)



	err = validateState(req, sess)
	if err != nil {
		return goth.User{}, err
	}

	fmt.Println("validateState", err)


	user, err := provider.FetchUser(sess)
	if err == nil {
		// user can be found with existing session data
		return user, err
	}

	fmt.Println("user", user)



	// get new token and retry fetch
	_, err = sess.Authorize(provider, req.URL.Query())
	if err != nil {
		return goth.User{}, err
	}

	fmt.Println("sess.Authorize", err)

	err = StoreInSession(providerName, sess.Marshal(), req, res)
	fmt.Println("StoreInSession", err)


	if err != nil {
		return goth.User{}, err
	}

	gu, err := provider.FetchUser(sess)

	fmt.Println("provider.FetchUser", gu, "     ", err)

	return gu, err
}
*/
// validateState ensures that the state token param from the original
// AuthURL matches the one included in the current (callback) request.
/*func validateState(req *http.Request, session interfaces.Session) error {
	rawAuthURL, err := session.GetAuthURL()
	if err != nil {
		return err
	}

	authURL, err := url.Parse(rawAuthURL)
	if err != nil {
		return err
	}

	originalState := authURL.Query().Get("state")
	if originalState != "" && (originalState != req.URL.Query().Get("state")) {
		return errors.New("state token mismatch")
	}

	return nil
}*/

// Logout invalidates a user session.
func Logout(res http.ResponseWriter, req *http.Request, session interfaces.Session, sid string) error {
	err := session.Get([]byte(sid), session)
	if err != nil {
		return err
	}

	err = session.Delete([]byte(sid))
	if err != nil {
		return errors.New("Could not delete user session ")
	}
	return nil
}

// GetProviderName is a function used to get the name of a provider
// for a given request. By default, this provider is fetched from
// the URL query string. If you provide it in a different way,
// assign your own function to this variable that returns the provider
// name for your request.
var GetProviderName = getProviderName

func getProviderName(req *http.Request, session interfaces.Session, sid string) (string, error) {

	// get all the used providers
	providers := goth.GetProviders()

	// loop over the used providers, if we already have a valid session for any provider (ie. user is already logged-in with a provider), then return that provider name
	for _, provider := range providers {
		p := provider.Name()
		session.Get([]byte(sid), p)

	}

	// try to get it from the url param "provider"
	if p := req.URL.Query().Get("provider"); p != "" {
		return p, nil
	}

	// try to get it from the url param ":provider"
	if p := req.URL.Query().Get(":provider"); p != "" {
		return p, nil
	}

	// try to get it from the context's value of "provider" key
	if p, ok := mux.Vars(req)["provider"]; ok {
		return p, nil
	}

	//  try to get it from the go-context's value of "provider" key
	if p, ok := req.Context().Value("provider").(string); ok {
		return p, nil
	}

	// if not found then return an empty string with the corresponding error
	return "", errors.New("you must select a provider")
}


// UnmarshalSession will unmarshal a JSON string into a session.
func  UnmarshalSessionVK(session interfaces.Session, data string) (interfaces.Session, error) {
	err := json.NewDecoder(strings.NewReader(data)).Decode(&session)
	return session, err
}

func GetFromSession(key string, req *http.Request, session interfaces.Session) (string, error) {
	session, _ := Store.Get(req, SessionName)
	value, err := getSessionValue(session, key)
	if err != nil {
		return "", errors.New("could not find a matching session for this request")
	}

	return value, nil
}

func getSessionValue(session *sessions.Session, key string) (string, error) {
	value := session.Values[key]
	if value == nil {
		return "", fmt.Errorf("could not find a matching session for this request")
	}

	rdata := strings.NewReader(value.(string))
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return "", err
	}
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(s), nil
}