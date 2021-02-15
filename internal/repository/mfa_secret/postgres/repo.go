package postgres

import "github.com/jinzhu/gorm"

type MfaSecretRepository struct {
	Db *gorm.DB
}

Create(ctx context.Context, userID uint64, mfaSecret *entity.MfaSecret) error
Get(ctx context.Context, userID uint64) (*entity.MfaSecret, error)
Update(ctx context.Context, userID uint64, mfaSecret *entity.MfaSecret) error
Delete(ctx context.Context, userID uint64) error