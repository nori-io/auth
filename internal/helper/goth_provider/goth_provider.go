package goth_provider

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/vk"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type GothProviderHelper struct{}

func (h GothProviderHelper) Use(provider *entity.SocialProvider) {
	var p goth.Provider
	switch provider.Name {
	case "vk":
		p=vk.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	}
	if p!=nil{
		goth.UseProviders(p)
	}
}

func (h GothProviderHelper) UseAll(providers []entity.SocialProvider) {
	for _,v:=range providers{
		h.Use(&v)
	}
}

