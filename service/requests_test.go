package service_test

import (
	"testing"

	rest "github.com/cheebo/gorest"
	"github.com/stretchr/testify/assert"

	"github.com/nori-io/authentication/service"
)

func TestSignUpRequest_ValidateMail(t *testing.T) {
	a := assert.New(t)
	r := service.SignUpRequest{Email: "a"}
	a.Equal(r.ValidateMail(), false)

	r = service.SignUpRequest{Email: "test@example.com"}
	a.Equal(r.ValidateMail(), true)

}

func TestSignUpRequest_ValidatePhone(t *testing.T) {
	a := assert.New(t)
	r := service.SignUpRequest{PhoneCountryCode: "a", PhoneNumber: "b"}
	errPhoneCountryCode, errPhoneNumber := r.ValidatePhone()
	a.NotEqual(errPhoneCountryCode, nil)
	a.NotEqual(errPhoneNumber, nil)

	r = service.SignUpRequest{PhoneCountryCode: "8", PhoneNumber: "b"}
	errPhoneCountryCode, errPhoneNumber = r.ValidatePhone()
	a.Equal(errPhoneCountryCode, nil)
	a.NotEqual(errPhoneNumber, nil)

	r = service.SignUpRequest{PhoneCountryCode: "a", PhoneNumber: "1"}
	errPhoneCountryCode, errPhoneNumber = r.ValidatePhone()
	a.NotEqual(errPhoneCountryCode, nil)
	a.Equal(errPhoneNumber, nil)

	r = service.SignUpRequest{PhoneCountryCode: "1", PhoneNumber: "1"}
	errPhoneCountryCode, errPhoneNumber = r.ValidatePhone()
	a.Equal(errPhoneCountryCode, nil)
	a.Equal(errPhoneNumber, nil)

}

func TestSignUpRequest_Validate(t *testing.T) {
	a := assert.New(t)
	r := service.SignUpRequest{Email: "test@example.com", Password: "pass"}
	errResponse := rest.ErrFieldResp{
		Meta: rest.ErrMeta{
			ErrCode: 400,
		},
		Fields: []rest.ErrField{},
	}
	a.Equal(r.Validate(), rest.ValidateResponse(errResponse, 400))

}

func TestSignUpRequest_ValidateMfaType(t *testing.T) {
	a := assert.New(t)
	r := service.SignUpRequest{MfaType: ""}
	a.Equal(r.ValidateMfaType(), nil)

	r = service.SignUpRequest{MfaType: "phone"}
	a.Equal(r.ValidateMfaType(), nil)

	r = service.SignUpRequest{MfaType: "otp"}
	a.Equal(r.ValidateMfaType(), nil)

	r = service.SignUpRequest{MfaType: "falseType"}
	a.NotEqual(r.ValidateMfaType(), nil)

}
