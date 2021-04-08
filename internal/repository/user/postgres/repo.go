package postgres

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/pkg/errors"

	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/pkg/transactor"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserRepository struct {
	Tx transactor.Transactor
}

func (r *UserRepository) Count(ctx context.Context) (uint64, error) {
	var count uint64
	if err := r.Tx.GetDB(ctx).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r *UserRepository) Create(ctx context.Context, e *entity.User) error {
	modelUser := NewModel(e)

	lastRecord := new(model)

	if err := r.Tx.GetDB(ctx).Create(&modelUser).Scan(&lastRecord).Error; err != nil {
		return errors.NewInternal(err)
	}

	return nil
}

func (r *UserRepository) FindById(ctx context.Context, id uint64) (*entity.User, error) {
	var (
		out = &model{}
		err error
	)
	err = r.Tx.GetDB(ctx).Where("id=?", id).First(out).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.Convert(), nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	out := &model{}

	err := r.Tx.GetDB(ctx).Where("email=?", email).First(&out).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.Convert(), nil
}

func (r *UserRepository) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	out := &model{}

	//@todo find by phone number and country code
	err := r.Tx.GetDB(ctx).Where("CONCAT(phone_number, phone_country_code)=?", phone).First(out).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.Convert(), nil
}

func (r *UserRepository) FindByFilter(ctx context.Context, filter repository.UserFilter) ([]entity.User, error) {
	var (
		models   []model
		entities []entity.User
	)
	q := r.Tx.GetDB(ctx).Offset(filter.Offset).Limit(filter.Limit)
	if filter.EmailPattern != nil {
		q = q.Where("email LIKE ?", filter.EmailPattern)
	}
	if filter.PhonePattern != nil {
		q = q.Where("CONCAT(phone_number, phone_country_code) LIKE ?", filter.PhonePattern)
	}

	if filter.UserStatus != nil {
		q = q.Where("status = ?", filter.UserStatus.Value())
	}

	err := q.Find(&models).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}
	for _, v := range models {
		entities = append(entities, *v.Convert())
	}
	return entities, nil
}

func (r *UserRepository) Update(ctx context.Context, e *entity.User) error {
	model := NewModel(e)
	if err := r.Tx.GetDB(ctx).Save(model).Error; err != nil {
		return errors.NewInternal(err)
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	if err := r.Tx.GetDB(ctx).Delete(&model{ID: id}).Error; err != nil {
		return errors.NewInternal(err)
	}
	return nil
}
