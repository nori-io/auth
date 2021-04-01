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

func (srv UserService) CreateUser(tx *gorm.DB, ctx context.Context, data service.SignUpData) (*entity.User, error) {
	user, err := srv.userRepository.FindByEmail(ctx, data.Email)
	if err != nil {
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

	if err := srv.userRepository.Create(tx, ctx, user); err != nil {
		return nil, errors.NewInternal(err)
	}
	return user, nil
}
