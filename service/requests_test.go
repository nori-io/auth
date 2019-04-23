package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nori-io/authorization/service"
)


func TestSignUpRequest_ValidateMail(t *testing.T) {
    a:=assert.New(t)
	r:=service.SignUpRequest{Email:"a"}
	a.Equal(r.ValidateMail(),false)

	r=service.SignUpRequest{Email:"test@mail.ru"}
	a.Equal(r.ValidateMail(),true)

}

func TestSignUpRequest_ValidatePhone(t *testing.T) {
	a:=assert.New(t)
	r:=service.SignUpRequest{PhoneCountryCode:"a",PhoneNumber:"b"}
	errPhoneCountryCode,errPhoneNumber:=r.ValidatePhone()
	a.NotEqual(errPhoneCountryCode,nil)
	a.NotEqual(errPhoneNumber,nil)

	r=service.SignUpRequest{PhoneCountryCode:"8",PhoneNumber:"b"}
	errPhoneCountryCode,errPhoneNumber=r.ValidatePhone()
	a.Equal(errPhoneCountryCode,nil)
	a.NotEqual(errPhoneNumber,nil)

	r=service.SignUpRequest{PhoneCountryCode:"a",PhoneNumber:"1"}
	errPhoneCountryCode,errPhoneNumber=r.ValidatePhone()
	a.NotEqual(errPhoneCountryCode,nil)
	a.Equal(errPhoneNumber,nil)

	r=service.SignUpRequest{PhoneCountryCode:"1",PhoneNumber:"1"}
	errPhoneCountryCode,errPhoneNumber=r.ValidatePhone()
	a.Equal(errPhoneCountryCode,nil)
	a.Equal(errPhoneNumber,nil)


}


func TestSignUpRequest_Validate(t *testing.T) {
	a := assert.New(t)
	r := service.SignUpRequest{Email: "test@mail.ru", Password: "pass"}
	a.Equal(r.Validate(), nil)

}

func TestSignUpRequest_ValidateMfaType(t *testing.T) {
	a:=assert.New(t)
	r:=service.SignUpRequest{MfaType:""}
	a.Equal(r.ValidateMfaType(),nil)

	r=service.SignUpRequest{MfaType:"phone"}
	a.Equal(r.ValidateMfaType(),nil)

	r=service.SignUpRequest{MfaType:"otp"}
	a.Equal(r.ValidateMfaType(),nil)

	r=service.SignUpRequest{MfaType:"falseType"}
	a.NotEqual(r.ValidateMfaType(),nil)


}

