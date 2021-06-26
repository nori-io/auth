package user_log

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

func (srv UserLogService) Create(ctx context.Context, data service.UserLogCreateData) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := srv.userLogRepository.Create(ctx, &entity.UserLog{
		UserID:    data.UserID,
		Action:    data.Action,
		SessionID: data.SessionID,
		Meta:      data.Meta,
		CreatedAt: data.CreatedAt,
	}); err != nil {
		return err
	}
	return nil
}
