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
}

func New(sessionInstance s.Session,
	userRepositoryInstance repository.UserRepository,
	mfaRecoveryCodeRepositoryInstance repository.MfaRecoveryCodeRepository) serv.AuthenticationService {
	return &service{
		session:                   sessionInstance,
		userRepository:            userRepositoryInstance,
		mfaRecoveryCodeRepository: mfaRecoveryCodeRepositoryInstance,
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
