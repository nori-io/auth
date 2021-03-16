package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeRepository interface {
	Delete(ctx context.Context, e *entity.MfaRecoveryCode) error
	Create(ctx context.Context, mfaRecoveryCode []entity.MfaRecoveryCode) error
}
