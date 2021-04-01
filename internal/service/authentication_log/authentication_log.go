package authentication_log

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/nori-plugins/authentication/pkg/enum/users_action"
	"github.com/nori-plugins/authentication/pkg/errors"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv AuthenticationLogService) CreateAuthenticationLog(tx *gorm.DB, ctx context.Context, user *entity.User) error {
	authenticationLog := &entity.AuthenticationLog{
		UserID: user.ID,
		Action: users_action.SignUp,
		//@todo заполнить метаданные айпи адресом и городом или чем-то ещё?
		CreatedAt: time.Now(),
	}

	if err := srv.authenticationLogRepository.Create(tx, ctx, authenticationLog); err != nil {
		return errors.NewInternal(err)
	}
	return nil
}
