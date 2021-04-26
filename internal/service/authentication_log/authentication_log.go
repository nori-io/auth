package authentication_log

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

func (srv AuthenticationLogService) Create(ctx context.Context, data service.CreateData) error {
	if err := srv.authenticationLogRepository.Create(ctx, &entity.AuthenticationLog{
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
