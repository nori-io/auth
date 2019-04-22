package service

import (
	"context"

	rest "github.com/cheebo/gorest"
	"github.com/cheebo/rand"
	"github.com/nori-io/nori-common/interfaces"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/authorization/service/database"
)

type Service interface {
	SignUp(ctx context.Context, req SignUpRequest) (resp *SignUpResponse)
	SignIn(ctx context.Context, req SignInRequest) (resp *SignInResponse)
	SignOut(ctx context.Context, req SignOutRequest) (resp *SignOutResponse)
	RecoveryCodes(ctx context.Context, req RecoveryCodesRequest) (resp *RecoveryCodesResponse)
}

type Config struct {
	Sub                          func() string
	Iss                          func() string
	UserType                     func() []interface{}
	UserTypeDefault              func() string
	UserRegistrationPhoneNumber  func() bool
	UserRegistrationEmailAddress func() bool
	UserMfaType                  func() string
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
	cache interfaces.Cache,
	mail interfaces.Mail,
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
	if (req.Email == "") && (req.PhoneNumber+req.PhoneCountryCode == "") {
		resp.Err = rest.ErrFieldResp{
			Meta: rest.ErrFieldRespMeta{
				ErrMessage: "Email and Phone's information is absent",
			},
		}
		return resp
	}

	if (req.PhoneNumber == "") || (req.PhoneCountryCode == "") {

	}

	if req.Email != "" {
		if modelAuth, err = s.db.Auth().FindByEmail(req.Email); err != nil {
			resp.Err = rest.ErrFieldResp{
				Meta: rest.ErrFieldRespMeta{
					ErrMessage: err.Error(),
				},
			}
			return resp
		}
		if modelAuth != nil && modelAuth.Id != 0 {
			errField.AddError("email", 400, "Email already exists.")
		}
	}
	if (req.PhoneCountryCode + req.PhoneCountryCode) != "" {
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

func (s *service) SignIn(ctx context.Context, req SignInRequest) (resp *SignInResponse) {
	resp = &SignInResponse{}

	modelFindEmail, errFindEmail := s.db.Auth().FindByEmail(req.Name)
	modelFindPhone, errFindPhone := s.db.Auth().FindByPhone(req.Name, "")
	if (errFindEmail != nil) && (errFindPhone != nil) {
		resp.Err = rest.ErrorInternal("Internal error")
		return resp
	}

	if (modelFindEmail.Id == 0) && (modelFindPhone.Id == 0) {
		resp.Err = rest.ErrorNotFound("User not found")
		return resp
	}

	var UserIdTemp uint64
	if modelFindEmail.Id != 0 {
		UserIdTemp = modelFindEmail.Id
		result, err := database.Authenticate([]byte(req.Password), modelFindEmail.Salt, modelFindEmail.Password)

		if (result == false) || (err != nil) {
			resp.Err = rest.ErrorNotFound("Uncorrect Password")
			return resp
		}

	}

	if modelFindPhone.Id != 0 {

		UserIdTemp = modelFindPhone.Id
		result, err := database.Authenticate([]byte(req.Password), modelFindPhone.Salt, modelFindPhone.Password)

		if (result == false) || (err != nil) {
			resp.Err = rest.ErrorNotFound("Uncorrect Password")
			return resp
		}

	}
	modelAuthenticationHistory := &database.AuthenticationHistoryModel{
		UserId: UserIdTemp,
	}

	err := s.db.AuthenticationHistory().Create(modelAuthenticationHistory)
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

func (s *service) RecoveryCodes(ctx context.Context, req RecoveryCodesRequest) (resp *RecoveryCodesResponse) {
	var modelMfa *database.MfaCodeModel
	modelMfa = &database.MfaCodeModel{
		UserId: req.UserId,
	}

	resp = &RecoveryCodesResponse{}

	codes, err := s.db.MfaCode().Create(modelMfa)
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
