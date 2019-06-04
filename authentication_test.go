package main_test

import (
	"testing"
)

func TestPlugin_Start(t *testing.T) {
	/*a := assert.New(t)

	plugin := auth.Plugin

	sDescr := mock.MatchedBy(func(key string) bool {
		return len(key) >= 0
	})

	// mock
	config := mocks.Config{}
	config.On("String", "jwt.sub", sDescr).Return(func() config2.String {
		return func() string {
			return "jwt.sub"
		}
	}())
	config.On("String", "jwt.iss", sDescr).Return(func() config2.String {
		return func() string {
			return "jwt.iss"
		}
	}())
	configManager := &mocks.Manager{}
	configManager.On("Register", mock.MatchedBy(func(m meta.Meta) bool {
		return true
	})).Return(&config)

	//
	auth := mocks2.Auth{}
	auth.On("")

	mail := &mocks2.Mail{}
	//
	registry := &mocks.Registry{}
	registry.On("Mail").Return(mail)

	mail.Send("")
	//
	ctx := context.Background()
	plugin.Init(ctx, configManager)
	err := plugin.Start(ctx, registry)
	a.Nil(err)*/
}
