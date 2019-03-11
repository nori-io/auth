package service

import (
	"context"

	rest "github.com/cheebo/gorest"
	"github.com/cheebo/rand"
	"github.com/nori-io/nori-common/interfaces"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/auth/service/database"
)

type Service interface {
	SignUp(ctx context.Context, req SignUpRequest) (resp *SignUpResponse)
	SignIn(ctx context.Context, req SignInRequest) (resp *SignInResponse)
	SignOut(ctx context.Context, req SignOutRequest) (resp *SignOutResponse)
}

type Config struct {
	Sub                          func() string
	Iss                          func() string
	UserType                     func() []interface{}
	UserTypeDefault              func() string
	UserRegistrationPhoneNumber  func() bool
	UserRegistrationEmailAddress func() bool
}

type service struct {
	auth    interfaces.Auth
	db      database.Database
	session interfaces.Session
	cfg     *Config
	log     *logrus.Logger
}

func NewService(
	auth interfaces.Auth,
	session interfaces.Session,
	cfg *Config,
	log *logrus.Logger,
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
	if (req.Email == "") && (req.Phone == "") {
		resp.Err = rest.ErrorInternal(err.Error())
		return resp
	}

	if req.Email != "" {
		if modelAuth, err = s.db.Auth().FindByEmail(req.Email); err != nil {
			resp.Err = rest.ErrorInternal(err.Error())
			return resp
		}
		if modelAuth != nil && modelAuth.Id != 0 {
			errField.AddError("email", 400, "Email already exists.")
		}
	}
	if req.Phone != "" {
		if modelAuth, err = s.db.Auth().FindByPhone(req.Phone); err != nil {
			resp.Err = rest.ErrorInternal(err.Error())
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
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
	}

	modelUsers = &database.UsersModel{
		Type: req.Type,
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
	resp.PhoneNumber = req.Phone

	return resp
}

func (s *service) SignIn(ctx context.Context, req SignInRequest) (resp *SignInResponse) {
	resp = &SignInResponse{}

	modelFindEmail, err := s.db.Auth().FindByEmail(req.Name)
	if err != nil {
		resp.Err = rest.ErrorInternal("Internal error")
		return resp
	}

	modelFindPhone, err := s.db.Auth().FindByPhone(req.Name)
	if err != nil {
		resp.Err = rest.ErrorInternal("Internal error")
		return resp
	}

	if (modelFindEmail == nil) && (modelFindPhone == nil) {
		resp.Err = rest.ErrorNotFound("User not found")
		return resp
	}

	if ((req.Password == modelFindEmail.Password) || (req.Password == modelFindPhone.Password)) == false {

		resp.Err = rest.ErrorNotFound("Uncorrect Password")
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

	sid := rand.RandomAlphaNum(32)

	token, err := s.auth.AccessToken(func(op interface{}) interface{} {
		key, ok := op.(string)
		if !ok || key == "" {
			return ""
		}
		switch key {
		case "raw":
			return map[string]string{
				"id":   string(UserIdTemp),
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

	s.session.Save([]byte(sid), interfaces.SessionActive, 0)

	resp.Id = uint64(UserIdTemp)
	resp.Token = token

	if modelFindEmail.Id != 0 {
		resp.User = *modelFindEmail
	}

	if modelFindPhone.Id != 0 {
		resp.User = *modelFindPhone
	}

	return resp
}

func (s *service) SignOut(ctx context.Context, req SignOutRequest) (resp *SignOutResponse) {
	resp = &SignOutResponse{}
	s.session.Delete(s.session.SessionId(ctx))
	return resp
}
