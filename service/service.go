package service

import (
	"context"
	"fmt"
	"reflect"
	"time"

	rest "github.com/cheebo/gorest"
	"github.com/cheebo/rand"
	"github.com/nori-io/nori-common/interfaces"
	"github.com/pkg/errors"

	"github.com/nori-io/authentication/service/database"
)

type Service interface {
	SignUp(ctx context.Context, req SignUpRequest) (resp *SignUpResponse)
	SignIn(ctx context.Context, req SignInRequest, parameters PluginParameters) (resp *SignInResponse)
	SignOut(ctx context.Context, req SignOutRequest) (resp *SignOutResponse)
	RecoveryCodes(ctx context.Context, req RecoveryCodesRequest) (resp *RecoveryCodesResponse)
}

type Config struct {
	Sub                            func() string
	Iss                            func() string
	UserType                       func() []interface{}
	UserTypeDefault                func() string
	UserRegistrationByPhoneNumber  func() bool
	UserRegistrationByEmailAddress func() bool
	UserMfaType                    func() string
}

type service struct {
	auth    interfaces.Auth
	db      database.Database
	session interfaces.Session
	cfg     *Config
	log     interfaces.Logger
}

type sessionData struct {
	name string
}

func NewService(
	auth interfaces.Auth,
	cache interfaces.Cache,
	mail interfaces.Mail,
	session interfaces.Session,
	cfg *Config,
	log interfaces.Logger,
	db database.Database,
) Service {
	return &service{
		auth:    auth,
		db:      db,
		session: session,
		cfg:     cfg,
		log:     log,
	}
}

func (s *service) SignUp(ctx context.Context, req SignUpRequest) (resp *SignUpResponse) {
	var err error
	var modelAuth *database.AuthModel
	var modelUsers *database.UsersModel
	resp = &SignUpResponse{}

	errField := rest.ErrFieldResp{
		Meta: rest.ErrFieldRespMeta{
			ErrCode: 400,
		},
	}
	if len(req.Email) != 0 {
		modelAuth, err = s.db.Auth().FindByEmail(req.Email)
	} else if len(req.PhoneCountryCode+req.PhoneCountryCode) != 0 {
		modelAuth, err = s.db.Auth().FindByEmail(req.Email)
	}

	if modelAuth != nil {
		errField.AddError("email", 400, "Email already exists.")
	}

	if len(req.PhoneCountryCode+req.PhoneNumber) != 0 {
		if modelAuth, err = s.db.Auth().FindByPhone(req.PhoneCountryCode, req.PhoneNumber); err != nil {
			resp.Err = rest.ErrFieldResp{
				Meta: rest.ErrFieldRespMeta{
					ErrMessage: err.Error(),
				},
			}
			return resp
		}
		if modelAuth != nil && modelAuth.Id != 0 {
			errField.AddError("phone", 400, "Phone number already exists.")
		}
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

	err = s.db.Users().Create(modelAuth, modelUsers)
	if err != nil {
		s.log.Error(err)
		resp.Err = rest.ErrFieldResp{
			Meta: rest.ErrFieldRespMeta{
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
	var model *database.AuthModel
	var err error

	if parameters.UserRegistrationByEmailAddress {
		model, err = s.db.Auth().FindByEmail(req.Name)
	} else {
		if parameters.UserRegistrationByPhoneNumber {
			model, err = s.db.Auth().FindByPhone(req.Name, "")
		}
	}

	if err != nil {
		resp.Err = rest.ErrorInternal("Database error")
		return resp
	}

	if model == nil {
		resp.Err = rest.ErrorNotFound("User not found")
		return resp
	}

	var userId uint64
	if model.Id != 0 {
		userId = model.Id
		result, err := database.VerifyPassword([]byte(req.Password), model.Salt, model.Password)

		if (!result) || (err != nil) {
			resp.Err = rest.ErrorNotFound("Uncorrect Password")
			return resp
		}

	}

	modelAuthenticationHistory := &database.AuthenticationHistoryModel{
		UserId: userId,
	}

	if model != nil {
		err = s.db.AuthenticationHistory().Create(modelAuthenticationHistory)
		if err != nil {
			s.log.Error(err)
			resp.Err = rest.ErrFieldResp{
				Meta: rest.ErrFieldRespMeta{
					ErrCode:    500,
					ErrMessage: err.Error(),
				},
			}
			return resp
		}
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
		resp.User = *model
	}

	return resp
}

func (s *service) SignOut(ctx context.Context, req SignOutRequest) (resp *SignOutResponse) {

	resp = &SignOutResponse{}

	value := ctx.Value("nori.auth.data")

	bar, err := InterfaceMap(value)
	if err != nil {
		panic(err)
	}
	var name string

	if val, ok := bar.(map[string]interface{})["raw"]; ok {
		if val2, ok2 := val.(map[string]interface{})["name"]; ok2 {
			name = fmt.Sprint(val2)
		}

	}

	/*tempData:=sessionData{}
	sessionId:=s.session.SessionId(ctx)

	err=s.session.Get(sessionId, &tempData)
	*/

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
	if modelFindEmail.Id != 0 {
		UserIdTemp = modelFindEmail.Id

	}

	if modelFindPhone.Id != 0 {

		UserIdTemp = modelFindPhone.Id

	}
	modelAuthenticationHistory := &database.AuthenticationHistoryModel{

		UserId: UserIdTemp,
	}

	if modelFindEmail.Id != 0 {
		modelAuthenticationHistory.SignOut = time.Now()
		err = s.db.AuthenticationHistory().Update(modelAuthenticationHistory)
	}
	if err != nil {
		s.log.Error(err)
		resp.Err = rest.ErrFieldResp{
			Meta: rest.ErrFieldRespMeta{
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
			Meta: rest.ErrFieldRespMeta{
				ErrCode:    500,
				ErrMessage: err.Error(),
			},
		}
		return resp
	}
	resp.Codes = codes
	return resp
}

/*func (s *service) MakeProfileEndpoint(ctx context.Context,req ProfileRequest)(resp *ProfileRequest){
	return resp
}*/

func InterfaceMap(i interface{}) (interface{}, error) {
	// Get type
	t := reflect.TypeOf(i)

	switch t.Kind() {
	case reflect.Map:
		// Get the value of the provided map
		v := reflect.ValueOf(i)

		// The "only" way of making a reflect.Type with interface{}
		it := reflect.TypeOf((*interface{})(nil)).Elem()

		// Create the map of the specific type. Key type is t.Key(), and element type is it
		m := reflect.MakeMap(reflect.MapOf(t.Key(), it))

		// Copy values to new map
		for _, mk := range v.MapKeys() {
			m.SetMapIndex(mk, v.MapIndex(mk))
		}

		return m.Interface(), nil

	}

	return nil, errors.New("Unsupported type")
}
