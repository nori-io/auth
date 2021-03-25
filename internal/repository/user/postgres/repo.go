package postgres

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserRepository struct {
	Db *gorm.DB
}

func (r *UserRepository) Count(ctx context.Context) (uint64, error) {
	var count uint64
	err := r.Db.Count(&count).Error
	return count, err
}

func (r *UserRepository) Create(tx *gorm.DB, ctx context.Context, e *entity.User) error {
	modelUser := NewModel(e)

	lastRecord := new(model)

	if err := tx.Create(modelUser).Scan(&lastRecord).Error; err != nil {
		tx.Rollback()
		return err
	}
	lastRecord.Convert()

	return nil
}

func (r *UserRepository) FindById(ctx context.Context, id uint64) (*entity.User, error) {
	var (
		out = &model{}
		e   error
	)
	e = r.Db.Where("id=?", id).First(out).Error

	return out.Convert(), e
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	out := &model{}

	err := r.Db.Where("email=?", email).First(out).Error

	return out.Convert(), err
}

func (r *UserRepository) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	out := &model{}

	//@todo find by phone number and country code
	err := r.Db.Where("CONCAT(phone_number, phone_country_code)=?", phone).First(out).Error

	return out.Convert(), err
}

func (r *UserRepository) FindByFilter(ctx context.Context, filter repository.UserFilter) ([]entity.User, error) {
	var (
		models   []model
		entities []entity.User
	)
	q := r.Db.Offset(filter.Offset).Limit(filter.Limit)
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
	if err != nil {
		return nil, err
	}
	for _, v := range models {
		entities = append(entities, *v.Convert())
	}
	return entities, err
}

func (r *UserRepository) Update(ctx context.Context, e *entity.User) error {
	model := NewModel(e)
	err := r.Db.Save(model).Error

	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	if err := r.Db.Delete(&model{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
