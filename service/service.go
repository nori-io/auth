package service

import (
	"context"
	"github.com/cheebo/gorest"
	"github.com/cheebo/rand"
	"github.com/sirupsen/logrus"
	"github.com/nori-io/auth/service/database"
	"github.com/nori-io/nori-common/interfaces"
)


type Service interface {
	SignUp(ctx context.Context, req SignUpRequest) (resp *SignUpResponse)
	Login(ctx context.Context, req LoginRequest) (resp *LogInResponse)
	Logout(ctx context.Context, req LogoutRequest) (resp *LogOutResponse)
}

type Config struct {
	Sub func() string
	Iss func() string
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
	var model *database.UsersModel
	resp = &SignUpResponse{}

	errField := rest.ErrFieldResp{
		Meta: rest.ErrFieldRespMeta{
			ErrCode: 400,
		},
	}

	if model, err = s.db.Users().FindByEmail(req.Email); err != nil {
		resp.Err = rest.ErrorInternal(err.Error())
		return resp
	}
	if model != nil && model.Id != 0 {
		errField.AddError("email", 400, "Email already exists.")
	}
	if errField.HasErrors() {
		resp.Err = errField
		return resp
	}

	model = &database.UsersModel{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	err = s.db.Users().Create(model)
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

	resp.Name = req.Name
	resp.Email = req.Email

	return resp
}

func (s *service) Login(ctx context.Context, req LoginRequest) (resp *LogInResponse) {
	resp = &LogInResponse{}

	model, err := s.db.Users().FindByEmail(req.Email)
	if err != nil {
		resp.Err = rest.ErrorInternal("Internal error")
		return resp
	}
	if model == nil {
		resp.Err = rest.ErrorNotFound("User not found")
		return resp
	}

	if req.Password != model.Password {
		resp.Err = rest.ErrorNotFound("User not found")
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
				"id":    string(model.Id),
				"email": model.Email,
				"name":  model.Name,
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

	resp.Id = uint64(model.Id)
	resp.Token = token
	resp.User = *model

	return resp
}

func (s *service) Logout(ctx context.Context, req LogoutRequest) (resp *LogOutResponse) {
	resp = &LogOutResponse{}
	s.session.Delete(s.session.SessionId(ctx))
	return resp
}