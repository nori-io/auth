package service_test

import (
	"testing"

	"github.com/nori-io/nori-common/mocks"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/authentication/service"
)

func TestService_SignUp(t *testing.T) {

	registryTest := mocks.Registry{}
	cfg := &service.Config{}
	auth, err := registryTest.Auth()
	if err != nil {
		t.Log(err)
	}

	cache, err := registryTest.Cache()
	if err != nil {
		t.Log(err)
	}

	mail, err := registryTest.Mail()
	if err != nil {
		t.Log(err)
	}

	session, err := registryTest.Session()
	if err != nil {
		t.Log(err)
	}

	serviceTest := service.NewService(auth, cache, cfg, nil, new(logrus.Logger), mail, session)

}
