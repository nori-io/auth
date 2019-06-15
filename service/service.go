package service

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"time"

	rest "github.com/cheebo/gorest"
	"github.com/cheebo/rand"
	"github.com/dgrijalva/jwt-go"
	"github.com/nori-io/nori-common/logger"
	"github.com/nori-io/nori-interfaces/interfaces"

	"github.com/nori-io/authentication/service/database"
)

type Service interface {
	SignUp(ctx context.Context, req SignUpRequest, parameters PluginParameters) (resp *SignUpResponse)
	SignIn(ctx context.Context, req SignInRequest, parameters PluginParameters) (resp *SignInResponse)
	SignOut(ctx context.Context, req SignOutRequest) (resp *SignOutResponse)
	RecoveryCodes(ctx context.Context, req RecoveryCodesRequest) (resp *RecoveryCodesResponse)
	SignInSocial(res http.ResponseWriter, req http.Request) (resp *SignInSocialResponse)
	SignOutSocial(res http.ResponseWriter, req http.Request) (resp *SignOutSocialResponse)

	/*SignInSocial(ctx context.Context, req http.Request, parameters PluginParameters) (resp *SignInSocialResponse)
	SignOutSocial(res http.ResponseWriter, req *http.Request)*/
}

type Config struct {
	Sub                                func() string
	Iss                                func() string
	UserType                           func() []interface{}
	UserTypeDefault                    func() string
	UserRegistrationByPhoneNumber      func() bool
	UserRegistrationByEmailAddress     func() bool
	UserMfaType                        func() string
	ActivationTimeForActivationMinutes func() uint
	ActivationCode                     func() bool
	Oath2ProvidersVKClientKey          func() string
	Oath2ProvidersVKClientSecret       func() string
	Oath2ProvidersVKRedirectUrl        func() string
	Oath2SessionSecret                 func() string
}

type service struct {
	auth    interfaces.Auth
	cache   interfaces.Cache
	cfg     *Config
	db      database.Database
	log     logger.Writer
	mail    interfaces.Mail
	session interfaces.Session
}

type sessionData struct {
	name string
}

func NewService(
	auth interfaces.Auth,
	cache interfaces.Cache,
	cfg *Config,
	db database.Database,
	log logger.Writer,
	mail interfaces.Mail,
	session interfaces.Session,
) Service {
	return &service{
		auth:    auth,
		cache:   cache,
		cfg:     cfg,
		db:      db,
		log:     log,
		mail:    mail,
		session: session,
	}
}

func (s *service) SignUp(ctx context.Context, req SignUpRequest, parameters PluginParameters) (resp *SignUpResponse) {

	var err error
	var modelAuth *database.AuthModel
	var modelUsers *database.UsersModel
	resp = &SignUpResponse{}

	errField := rest.ErrFieldResp{
		Meta: rest.ErrMeta{
			ErrCode: 0,
		},
	}
	if len(req.Email) != 0 {
		modelAuth, err = s.db.Auth().FindByEmail(req.Email)
	} else if len(req.PhoneCountryCode+req.PhoneCountryCode) != 0 {
		modelAuth, err = s.db.Auth().FindByPhone(req.PhoneCountryCode, req.PhoneNumber)
	}

	if modelAuth != nil && modelAuth.Id != 0 {
		resp.Email = req.Email
		resp.PhoneCountryCode = req.PhoneCountryCode
		resp.PhoneNumber = req.PhoneNumber
		errField.AddError("phone, email", 400, "User already exists.")
	}

	if err != nil {
		resp.Err = err
		resp.Email = req.Email
		resp.PhoneCountryCode = req.PhoneCountryCode
		resp.PhoneNumber = req.PhoneNumber

		return resp
	}

	if errField.HasErrors() {

		resp.Err = errField
		return resp
	}

	modelAuth = &database.AuthModel{
		Email:            req.Email,
		Password:         []byte(req.Password),
		PhoneCountryCode: req.PhoneCountryCode,
		PhoneNumber:      req.PhoneNumber,
	}

	modelUsers = &database.UsersModel{
		Type:     req.Type,
		Mfa_type: req.MfaType,
	}

	if parameters.ActivationCode {
		modelUsers.Status_account = "locked"
	} else {
		modelUsers.Status_account = "active"
	}
	err = s.db.Users().Create(modelAuth, modelUsers)
	if err != nil {
		s.log.Error(err)
		resp.Err = rest.ErrFieldResp{
			Meta: rest.ErrMeta{
				ErrCode:    500,
				ErrMessage: err.Error(),
			},
		}

		return resp
	}

	resp.Email = req.Email
	resp.PhoneCountryCode = req.PhoneCountryCode
	resp.PhoneNumber = req.PhoneNumber

	return resp
}

func (s *service) SignIn(ctx context.Context, req SignInRequest, parameters PluginParameters) (resp *SignInResponse) {
	resp = &SignInResponse{}
	var model, modelFindByEmail, modelFindByPhone *database.AuthModel
	var errFindByEmail, errFindPhone error

	if parameters.UserRegistrationByEmailAddress {

		modelFindByEmail, errFindByEmail = s.db.Auth().FindByEmail(req.Name)
		if errFindByEmail != nil {
			resp.User.UserName = req.Name
			resp.Err = errFindByEmail
			return resp
		}
		if modelFindByEmail.Id != 0 {
			model = modelFindByEmail
		}
	}

	if parameters.UserRegistrationByPhoneNumber {
		modelFindByPhone, errFindPhone = s.db.Auth().FindByPhone(req.Name, "")
		if errFindPhone != nil {
			resp.User.UserName = req.Name
			resp.Err = errFindPhone
			return resp
		}
		if modelFindByPhone.Id != 0 {
			model = modelFindByPhone
		}
	}

	if model == nil {
		resp.User.UserName = req.Name
		resp.Err = rest.ErrResp{Meta: rest.ErrMeta{ErrMessage: "User not found", ErrCode: 0}}
		return resp
	}
	var userId uint64
	userId = model.Id
	result, err := database.VerifyPassword([]byte(req.Password), model.Salt, model.Password)

	if (!result) || (err != nil) {
		resp.Id = userId
		resp.User.UserName = req.Name
		resp.Err = rest.ErrResp{Meta: rest.ErrMeta{ErrMessage: "Incorrect Password", ErrCode: 0}}

		return resp
	}

	modelAuthenticationHistory := &database.AuthenticationHistoryModel{
		UserId: userId,
	}

	err = s.db.AuthenticationHistory().Create(modelAuthenticationHistory)
	if err != nil {
		s.log.Error(err)
		resp.User.UserName = req.Name
		resp.Err = rest.ErrFieldResp{
			Meta: rest.ErrMeta{
				ErrCode:    500,
				ErrMessage: err.Error(),
			},
		}
		return resp
	}

	sid := rand.RandomAlphaNum(32)

	token, err := s.auth.AccessToken(func(op interface{}) interface{} {
		key, ok := op.(string)
		if !ok || key == "" {
			return ""
		}
		switch key {
		case "raw":
			return map[string]string{
				"id":   string(userId),
				"name": req.Name,
			}
		case "jti":
			return sid
		case "sub":
			return s.cfg.Sub()
		case "iss":
			return s.cfg.Iss()
		default:
			return ""
		}
	})

	if err != nil {
		resp.Err = rest.ErrorInternal(err.Error())
		return resp
	}
	s.session.Save([]byte(sid), sessionData{name: req.Name}, 0)

	resp.Id = uint64(userId)
	resp.Token = token

	if model.Id != 0 {
		resp.User = UserResponse{UserName: req.Name}
	}

	return resp
}

func (s *service) SignOut(ctx context.Context, req SignOutRequest) (resp *SignOutResponse) {

	resp = &SignOutResponse{}

	value := ctx.Value("nori.auth.data")

	var name string

	if val, ok := value.(jwt.MapClaims)["raw"]; ok {
		reflect.TypeOf(val)
		if val2, ok2 := val.(map[string]interface{})["name"]; ok2 {
			name = fmt.Sprint(val2)
		}

	}

	req = SignOutRequest{}
	modelFindEmail, errFindEmail := s.db.Auth().FindByEmail(name)
	modelFindPhone, errFindPhone := s.db.Auth().FindByPhone(name, "")
	if (errFindEmail != nil) && (errFindPhone != nil) {
		resp.Err = rest.ErrorInternal("Internal error")
		return resp
	}

	if (modelFindEmail == nil) && (modelFindPhone == nil) {
		resp.Err = rest.ErrorNotFound("User not found")
		return resp
	}

	var UserIdTemp uint64
	if (modelFindEmail != nil) && (modelFindEmail.Id != 0) {
		UserIdTemp = modelFindEmail.Id

	}

	if (modelFindPhone != nil) && (modelFindPhone.Id != 0) {

		UserIdTemp = modelFindPhone.Id

	}
	modelAuthenticationHistory := &database.AuthenticationHistoryModel{

		UserId: UserIdTemp,
	}
	var err error
	if modelFindEmail.Id != 0 {
		modelAuthenticationHistory.SignOut = time.Now()
		err = s.db.AuthenticationHistory().Update(modelAuthenticationHistory)
	}
	if err != nil {
		s.log.Error(err)
		resp.Err = rest.ErrFieldResp{
			Meta: rest.ErrMeta{
				ErrCode:    500,
				ErrMessage: err.Error(),
			},
		}
		return resp
	}

	s.session.Delete(s.session.SessionId(ctx))

	return resp
}

func (s *service) RecoveryCodes(ctx context.Context, req RecoveryCodesRequest) (resp *RecoveryCodesResponse) {
	var modelMfa *database.MfaRecoveryCodesModel
	modelMfa = &database.MfaRecoveryCodesModel{
		UserId: req.UserId,
	}

	resp = &RecoveryCodesResponse{}

	codes, err := s.db.MfaRecoveryCodes().Create(modelMfa)
	if err != nil {
		s.log.Error(err)
		resp.Err = rest.ErrFieldResp{
			Meta: rest.ErrMeta{
				ErrCode:    500,
				ErrMessage: err.Error(),
			},
		}
		return resp
	}
	resp.Codes = codes
	return resp
}

func (s *service) SignInSocial(res http.ResponseWriter, req http.Request) (resp *SignInSocialResponse) {

	// try to get the user without re-authenticating
	gothUser, err := CompleteUserAuth(res, &req, s.session)
	fmt.Println("err is", err)
	if err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, gothUser)
	} else {
		BeginAuthHandler(res, &req, s.session)
		return nil
	}

	s.session.Save([]byte(gothUser.AccessToken), sessionData{name: gothUser.Email}, 0)

	var sd sessionData
	fmt.Println("session.Get", s.session.Get([]byte("zT6hfj3DshfF4ewehgwLsd91412dsW4F"), &sd))
	fmt.Println("sd.name", sd.name)

	return resp
}

func (s *service) SignOutSocial(res http.ResponseWriter, req http.Request) (resp *SignOutSocialResponse) {
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
	return nil
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
