package postgres

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (r *AuthRepository) Create(ctx context.Context, e *entity.User) error {
	return nil
}

func (r *AuthRepository) Get(ctx context.Context, id uint64) (*entity.User, error) {
	return nil, nil
}

func (r *AuthRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

func (r *AuthRepository) GetAll(ctx context.Context, offset uint64, limit uint64) ([]entity.User, error) {
	return nil, nil
}

func (r *AuthRepository) Update(ctx context.Context, e *entity.User) error {
	return nil
}

func (r *AuthRepository) Delete(ctx context.Context, id uint64) error {
	return nil
}
