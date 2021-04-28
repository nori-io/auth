package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaTotpRepository interface {
	Create(ctx context.Context, mfaTotp *entity.MfaTotp) error
	Update(ctx context.Context, mfaTotp *entity.MfaTotp) error
	Delete(ctx context.Context, userID uint64) error
}
