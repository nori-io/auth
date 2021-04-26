package entity

import (
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/social_provider_status"
)

type SocialProvider struct {
	ID          uint64
	AppID       string
	AppSecret   string
	Name        string
	Logo        string
	RedirectUrl string
	CallBackUrl string
	TokenUrl    string
	Scopes string
	Status      social_provider_status.Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
