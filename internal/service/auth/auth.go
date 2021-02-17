package auth

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	s "github.com/nori-io/interfaces/nori/session"
	serv "github.com/nori-plugins/authentication/internal/domain/service"
)

type service struct {
	session                   s.Session
	userRepository            repository.UserRepository
	mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	mfaSecretRepository       repository.MfaSecretRepository
	configData                configData
}

type configData struct {
	Issuer string
}

func New(sessionInstance s.Session,
	userRepositoryInstance repository.UserRepository,
	mfaRecoveryCodeRepositoryInstance repository.MfaRecoveryCodeRepository,
	mfaSecretRepositoryInstance repository.MfaSecretRepository,
	configData configData) serv.AuthenticationService {
	return &service{
		configData:                configData,
		session:                   sessionInstance,
		userRepository:            userRepositoryInstance,
		mfaRecoveryCodeRepository: mfaRecoveryCodeRepositoryInstance,
		mfaSecretRepository:       mfaRecoveryCodeRepositoryInstance,
	}
}

func (srv *service) SignUp(ctx context.Context, data serv.SignUpData) (*entity.User, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	var user *entity.User

	user = &entity.User{
		Email:    data.Email,
		Password: data.Password,
	}

	if err := srv.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (srv *service) SignIn(ctx context.Context, data serv.SignInData) (*entity.Session, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	user := &entity.User{
		Email:    data.Email,
		Password: data.Password,
	}

	var err error
	user, err = srv.userRepository.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	sid, err := srv.getToken()
	if err != nil {
		return nil, err
	}
	return &entity.Session{SessionKey: sid}, nil
}

func (srv *service) SignOut(ctx context.Context, data *entity.Session) error {
	err := srv.session.Delete([]byte(data.SessionKey))
	return err
}

func (srv *service) GetMfaRecoveryCodes(ctx context.Context, data *entity.Session) ([]entity.MfaRecoveryCode, error) {
	var codes []entity.MfaRecoveryCode
	var err error
	var mfaRecoveryCode *entity.MfaRecoveryCode
	//@todo read count of symbols from config
	//@todo read pattenn from config
	//@todo read symbol sequence from config
	//@todo generating of specify sequence
	//@todo нужна ли максимальная длина, или указать всё в паттерне?
	for i := 0; i < 10; i++ {
		sid := make([]byte, 32)

		if _, err := rand.Read(sid); err != nil {
			return nil, err
		}
		mfaRecoveryCode = &entity.MfaRecoveryCode{
			UserID:    data.UserID,
			Code:      string(sid),
			CreatedAt: time.Now(),
		}
		err = srv.mfaRecoveryCodeRepository.Create(ctx, data.UserID, mfaRecoveryCode)
		if err != nil {
			break
		}
		codes = append(codes, *mfaRecoveryCode)
	}
	return codes, err
}

func (srv *service) PostSecret(
	ctx context.Context, data *serv.SecretData, session entity.Session) (
	login string, issuer string, err error) {
	if err := data.Validate(); err != nil {
		return "", "", err
	}

	var mfaSecret *entity.MfaSecret

	mfaSecret = &entity.MfaSecret{
		UserID: session.UserID,
		Secret: data.Secret,
	}

	if err := srv.mfaSecretRepository.Create(ctx, mfaSecret); err != nil {
		return "", "", err
	}

	userData, err := srv.userRepository.Get(ctx, session.UserID)
	if err != nil {
		return "", "", err
	}

	if userData.Email != "" {
		login = userData.Email
	} else {
		login = userData.PhoneCountryCode + userData.PhoneNumber
	}
	return login, srv.configData.Issuer, nil
}

func (srv *service) getToken() ([]byte, error) {
	sid := make([]byte, 32)

	if _, err := rand.Read(sid); err != nil {
		return nil, err
	}
	if err := srv.session.Get(sid, s.SessionActive); err != nil {
		srv.session.Save(sid, s.SessionActive, 0)
		return sid, nil
	}
	return srv.getToken()
}
