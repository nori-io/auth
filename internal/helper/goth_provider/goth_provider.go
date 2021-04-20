package goth_provider

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/amazon"
	"github.com/markbates/goth/providers/apple"
	"github.com/markbates/goth/providers/auth0"
	"github.com/markbates/goth/providers/azuread"
	"github.com/markbates/goth/providers/battlenet"
	"github.com/markbates/goth/providers/bitbucket"
	"github.com/markbates/goth/providers/box"
	"github.com/markbates/goth/providers/dailymotion"
	"github.com/markbates/goth/providers/deezer"
	"github.com/markbates/goth/providers/digitalocean"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/dropbox"
	"github.com/markbates/goth/providers/eveonline"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/fitbit"
	"github.com/markbates/goth/providers/gitea"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/gitlab"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/gplus"
	"github.com/markbates/goth/providers/heroku"
	"github.com/markbates/goth/providers/instagram"
	"github.com/markbates/goth/providers/intercom"
	"github.com/markbates/goth/providers/kakao"
	"github.com/markbates/goth/providers/lastfm"
	"github.com/markbates/goth/providers/line"
	"github.com/markbates/goth/providers/linkedin"
	"github.com/markbates/goth/providers/mastodon"
	"github.com/markbates/goth/providers/meetup"
	"github.com/markbates/goth/providers/microsoftonline"
	"github.com/markbates/goth/providers/naver"
	"github.com/markbates/goth/providers/nextcloud"
	"github.com/markbates/goth/providers/okta"
	"github.com/markbates/goth/providers/onedrive"
	"github.com/markbates/goth/providers/openidConnect"
	"github.com/markbates/goth/providers/paypal"
	"github.com/markbates/goth/providers/salesforce"
	"github.com/markbates/goth/providers/seatalk"
	"github.com/markbates/goth/providers/shopify"
	"github.com/markbates/goth/providers/slack"
	"github.com/markbates/goth/providers/soundcloud"
	"github.com/markbates/goth/providers/spotify"
	"github.com/markbates/goth/providers/steam"
	"github.com/markbates/goth/providers/strava"
	"github.com/markbates/goth/providers/stripe"
	"github.com/markbates/goth/providers/twitch"
	"github.com/markbates/goth/providers/twitter"
	"github.com/markbates/goth/providers/typetalk"
	"github.com/markbates/goth/providers/uber"
	"github.com/markbates/goth/providers/vk"
	"github.com/markbates/goth/providers/wepay"
	"github.com/markbates/goth/providers/xero"
	"github.com/markbates/goth/providers/yahoo"
	"github.com/markbates/goth/providers/yammer"
	"github.com/markbates/goth/providers/yandex"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type GothProviderHelper struct{}

func (h GothProviderHelper) Use(provider *entity.SocialProvider) {
	var p goth.Provider
	switch provider.Name {
	case "amazon":
		p=amazon.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "bitbucket":
		p=bitbucket.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "box":
		p=box.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "dailymotion":
		p=dailymotion.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "deezer":
		p=deezer.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "digitalocean":
		p=digitalocean.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "discord":
		p=discord.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "dropbox":
		p=dropbox.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "eveonline":
		p=eveonline.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "facebook":
		p=facebook.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "fitbit":
		p=fitbit.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "gitea":
		p=gitea.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "github":
		p=github.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "gitlab":
		p=gitlab.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "google":
		p=google.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "gplus":
		p=gplus.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "shopify":
		p=shopify.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "soundcloud":
		p=soundcloud.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "spotify":
		p=spotify.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "steam":
		p=steam.New(provider.AppID, provider.RedirectUrl)
	case "stripe":
		p=stripe.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "twitch":
		p=twitch.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "uber":
		p=uber.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "wepay":
		p=wepay.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "yahoo":
		p=yahoo.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "yammer":
		p=yammer.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "heroku":
		p=heroku.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "instagram":
		p=instagram.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "intercom":
		p=intercom.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "kakao":
		p=kakao.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "lastfm":
		p=lastfm.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "linkedin":
		p=linkedin.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "line":
		p=line.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "onedrive":
		p=onedrive.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "azuread":
		p=azuread.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "microsoftonline":
		p=microsoftonline.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "battlenet":
		p=battlenet.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "paypal":
		p=paypal.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "twitter":
		p=twitter.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "salesforce":
		p=salesforce.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "typetalk":
		p=typetalk.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "slack":
		p=slack.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "meetup":
		p=meetup.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "auth0":
		p=auth0.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "openid-connect":
		openidConnect.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "xero":
		xero.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "vk":
		p = vk.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "naver":
		p=naver.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "yandex":
		p=yandex.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "nextcloud":
		p=nextcloud.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "seatalk":
		p=seatalk.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "apple":
		p=apple.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "strava":
		p=strava.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "okta":
		p=okta.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
	case "mastodon":
		p=mastodon.New(provider.AppID, provider.AppSecret, provider.RedirectUrl)
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

