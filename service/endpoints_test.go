package service_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"testingh

	"github.com/nori-io/nori-common/mocks"

	"github.com/nori-io/auth/service"
)

func TestMakeRecoveryCodesEndpoint(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	testService,w,w2,w3:=service.NewService((new(mocks.Auth),new (mocks.Session)), new (mocks.Config),new(logrus.Logger),mockDatabase)
	fmt.Print(testService,w,w2,w3)
	//service.MakeRecoveryCodesEndpoint()
}
