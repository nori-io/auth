package user

import (
	"context"
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/users_action"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/enum/hash_algorithm"
	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"
	"github.com/nori-plugins/authentication/pkg/enum/users_status"
	"github.com/nori-plugins/authentication/pkg/enum/users_type"
)

func (srv UserService) Create(ctx context.Context, data service.UserCreateData) (*entity.User, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	user, err := srv.GetByEmail(ctx, service.GetByEmailData{Email: data.Email})
	if err != nil && err != errors2.UserNotFound {
		return nil, err
	}
	if user != nil {
		return nil, errors2.EmailAlreadyTaken
	}

	password, err := srv.securityHelper.GenerateHash(data.Password)
	if err != nil {
		return nil, err
	}
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

	if err := srv.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (srv UserService) UpdatePassword(ctx context.Context, data service.UserUpdatePasswordData) error {
	if err := data.Validate(); err != nil {
		return err
	}

	password, err := srv.securityHelper.GenerateHash(data.Password)
	if err != nil {
		return err
	}
	if err := srv.userRepository.Update(ctx, &entity.User{
		ID:        data.UserID,
		Password:  string(password),
		UpdatedAt: time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

func (srv UserService) UpdateMfaStatus(ctx context.Context, data service.UserUpdateMfaStatusData) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := srv.userRepository.Update(ctx, &entity.User{
		ID:        data.UserID,
		MfaType:   data.MfaType,
		UpdatedAt: time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

func (srv UserService) UpdateUserStatus(ctx context.Context, data service.UserUpdateStatusData) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := srv.transactor.Transact(ctx, func(tx context.Context) error {
		if err := srv.userRepository.Update(ctx, &entity.User{
			ID:        data.UserID,
			Status:    data.Status,
			UpdatedAt: time.Now(),
		}); err != nil {
			return err
		}
		if err := srv.userLogService.Create(ctx, service.UserLogCreateData{
			UserID:    data.UserID,
			Action:    users_action.UserStatusChanged,
			Meta:      "",
			CreatedAt: time.Now(),
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (srv UserService) GetByEmail(ctx context.Context, data service.GetByEmailData) (*entity.User, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	user, err := srv.userRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors2.UserNotFound
	}
	return user, nil
}

func (srv UserService) GetByID(ctx context.Context, data service.GetByIdData) (*entity.User, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	user, err := srv.userRepository.FindByID(ctx, data.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors2.UserNotFound
	}
	return user, nil
}

func (srv UserService) GetAll(ctx context.Context, filter service.UserFilter) ([]entity.User, error) {
	users, err := srv.userRepository.FindByFilter(ctx, repository.UserFilter{
		EmailPattern: filter.EmailPattern,
		PhonePattern: filter.PhonePattern,
		UserStatus:   filter.UserStatus,
		Offset:       filter.Offset,
		Limit:        filter.Limit,
	})
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, errors2.UserNotFound
	}
	return users, nil
}
