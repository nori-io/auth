package reset_password

import (
	"context"
	"fmt"
	"time"

	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"

	"github.com/nori-plugins/authentication/internal/domain/entity"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

func (srv ResetPasswordService) RequestResetPasswordEmail(ctx context.Context, data service.RequestResetPasswordEmailData) error {
	if err := data.Validate(); err != nil {
		return err
	}
	user, err := srv.userService.GetByEmail(ctx, service.GetByEmailData{Email: data.Email})
	if err != nil {
		return err
	}
	if user == nil {
		return nil
	}

	token, err := srv.securityHelper.GenerateToken(32)
	if err != nil {
		return err
	}
	//@todo send to email
	fmt.Println(token)
	srv.resetPasswordRepository.Create(ctx, &entity.ResetPassword{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Duration(srv.config.EmailActivationCodeTTLSeconds()) * time.Second),
	})
	return nil
}

func (srv ResetPasswordService) SetNewPasswordByRestorePasswordEmailToken(ctx context.Context, data service.SetNewPasswordByRestorePasswordEmailTokenData) error {
	if err := data.Validate(); err != nil {
		return err
	}

	resetPassword, err := srv.resetPasswordRepository.FindByToken(ctx, data.Token)
	if err != nil {
		return err
	}
	if resetPassword == nil {
		return errors2.TokenNotFound
	}
	if resetPassword.ExpiresAt.Before(time.Now()) {
		return errors2.TokenNotFound
	}

	if err := srv.userService.UpdatePassword(ctx, service.UserUpdatePasswordData{
		UserID:   resetPassword.UserID,
		Password: data.Password,
	}); err != nil {
		return err
	}

	return nil
}
