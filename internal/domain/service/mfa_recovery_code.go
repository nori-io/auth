package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeService interface {
	GetMfaRecoveryCodes(ctx context.Context, data *entity.Session) ([]entity.MfaRecoveryCode, error)
	GetByUserId(ctx context.Context, userID uint64, code string) (*entity.MfaRecoveryCode, error)
	Apply(ctx context.Context, userID uint64, code string) error
}
