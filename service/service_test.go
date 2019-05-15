package service_test

import (
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cheebo/rand"
	"github.com/nori-io/nori-common/mocks"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/nori-io/authentication/service"
	"github.com/nori-io/authentication/service/database"
)


func TestService_SignUp(t *testing.T) {

	auth := &mocks.Auth{}

	cache := &mocks.Cache{}
	cfg := &service.Config{}
	mockDatabase, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db := database.DB(mockDatabase, logrus.New())
	mail:= &mocks.Mail{}
	session:= &mocks.Session{}


	serviceTest := service.NewService(auth, cache, cfg, db, new(logrus.Logger), mail, session)

	SignUpRequest:=service.SignUpRequest{Email:"test@mail.ru", Password:"pass", Type:"vendor"}

	serviceTest.SignUp(context.Background(), SignUpRequest)

	code := rand.RandomAlphaNum(65)
	salt, codeForSend, err := getActivationCode([]byte(code))
	if err == nil {
		//activationRequest:=service.ActivationCodeRequest{ActivationCode:codeForSend}
		activate([]byte(code),salt, codeForSend)
		mail.Send(activate([]byte(code),salt, codeForSend))
	}


}

func getActivationCode(codeMessage []byte) ([]byte, []byte, error) {
	salt := rand.RandomAlphaNum(65)

	codeForSend, err := database.Hash(codeMessage, []byte(salt))

	return []byte(salt), codeForSend, err
}

func activate(code []byte, salt []byte, codeForSend []byte ) string {
   result, _ := database.VerifyPassword([]byte(code), salt, codeForSend)
	b := strconv.FormatBool(result)

	return b
}