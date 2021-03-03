package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeService interface {
	GetMfaRecoveryCodes(ctx context.Context, data *entity.Session) ([]entity.MfaRecoveryCode, error)
}
