package service

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SocialProvider interface {
	GetAllActive(ctx context.Context) ([]entity.SocialProvider, error)
	GetByName(ctx context.Context, data GetByNameData) (*entity.SocialProvider, error)
}

type GetByNameData struct {
	Name string
}

func (d GetByNameData) Validate() error {
	return v.Errors{
		"name": v.Validate(d.Name, v.Required, v.Length(2, 254)),
	}.Filter()
}
