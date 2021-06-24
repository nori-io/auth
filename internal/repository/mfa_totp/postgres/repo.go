package postgres

import (
	"context"

	"gorm.io/gorm"

	"github.com/nori-plugins/authentication/pkg/errors"

	"github.com/nori-plugins/authentication/pkg/transactor"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaTotpRepository struct {
	Tx transactor.Transactor
}

func (r *MfaTotpRepository) Create(ctx context.Context, e *entity.MfaTotp) error {
	m := newModel(e)
	if err := r.Tx.GetDB(ctx).Create(m).Error; err != nil {
		return errors.NewInternal(err)
	}
	*e = *m.convert()

	return nil
}

func (r *MfaTotpRepository) Update(ctx context.Context, e *entity.MfaTotp) error {
	m := newModel(e)
	if err := r.Tx.GetDB(ctx).Save(m).Error; err != nil {
		return errors.NewInternal(err)
	}
	*e = *m.convert()

	return nil
}

func (r *MfaTotpRepository) Delete(ctx context.Context, userID uint64) error {
	if err := r.Tx.GetDB(ctx).Delete(&model{UserID: userID}).Error; err != nil {
		return errors.NewInternal(err)
	}
	return nil
}

func (r *MfaTotpRepository) FindByUserId(ctx context.Context, userID uint64) (*entity.MfaTotp, error) {
	out := &model{}

	err := r.Tx.GetDB(ctx).Where("user_id=?", userID).First(out).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}
	return out.convert(), nil
}
