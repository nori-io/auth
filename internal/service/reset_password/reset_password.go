package reset_password

import (
	"context"
	"fmt"
	"time"

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
		ExpiresAt: time.Now(),
	})
	return nil
}

func (srv ResetPasswordService) SetNewPasswordByRestorePasswordEmailToken(ctx context.Context, password string) error {
	panic("implement me")
}
