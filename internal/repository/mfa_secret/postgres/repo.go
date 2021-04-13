package postgres

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/pkg/errors"

	"github.com/nori-plugins/authentication/pkg/transactor"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaSecretRepository struct {
	Tx transactor.Transactor
}

func (r *MfaSecretRepository) Create(ctx context.Context, e *entity.MfaSecret) error {
	modelMfaSecret := newModel(e)

	lastRecord := new(model)

	if err := r.Tx.GetDB(ctx).Create(modelMfaSecret).Scan(&lastRecord).Error; err != nil {
		return errors.NewInternal(err)
	}
	lastRecord.convert()

	return nil
}

func (r *MfaSecretRepository) Get(ctx context.Context, userID uint64) (*entity.MfaSecret, error) {
	out := &model{}
	err := r.Tx.GetDB(ctx).Where("id=?", userID).First(out).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.convert(), nil
}

func (r *MfaSecretRepository) Update(ctx context.Context, userID uint64, e *entity.MfaSecret) error {
	model := newModel(e)
	if err := r.Tx.GetDB(ctx).Save(model).Error; err != nil {
		return errors.NewInternal(err)
	}

	return nil
}

func (r *MfaSecretRepository) Delete(ctx context.Context, userID uint64) error {
	if err := r.Tx.GetDB(ctx).Delete(&model{UserID: userID}).Error; err != nil {
		return errors.NewInternal(err)
	}
	return nil
}
