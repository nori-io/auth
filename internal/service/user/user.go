package user

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/enum/hash_algorithm"
	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"
	"github.com/nori-plugins/authentication/pkg/enum/users_status"
	"github.com/nori-plugins/authentication/pkg/enum/users_type"
	"github.com/nori-plugins/authentication/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (srv UserService) Create(ctx context.Context, data service.UserCreateData) (*entity.User, error) {
	user, err := srv.userRepository.FindByEmail(ctx, data.Email)
	if (err != nil) && (err != gorm.ErrRecordNotFound) {
		return nil, errors.NewInternal(err)
	}
	if user != nil {
		return nil, errors2.DuplicateUser
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), srv.config.PasswordBcryptCost())

	//@todo заполнить оставшиеся поля по мере разработки нового функционала
	user = &entity.User{
		Status:          users_status.Active,
		UserType:        users_type.User,
		MfaType:         mfa_type.None,
		Email:           data.Email,
		Password:        string(password),
		HashAlgorithm:   hash_algorithm.Bcrypt,
		IsEmailVerified: srv.config.EmailVerification(),
		CreatedAt:       time.Now(),
	}

	if err := srv.transactor.Transact(ctx, func(tx context.Context) error {
		if err := srv.userRepository.Create(tx, user); err != nil {
			return errors.NewInternal(err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return user, nil
}
