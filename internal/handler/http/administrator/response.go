package administrator

import (
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserResponse struct {
	ID      uint64    `json:"name"`
	Status  string    `json:"status"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Created time.Time `json:"created"`
}

func convertAll(entities []entity.User) []UserResponse {
	socialProviders := make([]UserResponse, 0)
	for _, v := range entities {
		socialProviders = append(socialProviders, convert(v))
	}
	return socialProviders
}

func convert(e entity.User) UserResponse {
	return UserResponse{
		ID:      e.ID,
		Status:  string(e.Status),
		Email:   e.Email,
		Phone:   e.PhoneCountryCode + e.PhoneNumber,
		Created: e.CreatedAt,
	}
}
