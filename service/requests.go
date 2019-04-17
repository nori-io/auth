package service

import (
	"github.com/asaskevich/govalidator"
	"github.com/cheebo/gorest"
)

// SignUp Request
type SignUpRequest struct {
	Email            string `json:"email" validate:"email"`
	PhoneCountryCode string `json:"phone_country_code"`
	PhoneNumber      string `json:"phone_number"`
	Password         string `json:"password" validate:"password"`
	Type             string `json:"user_type" validate:"user_type"`
	Meta             string `json:"meta" validate:"meta"`
	MfaType          string `json:mfa_type validate:"mfa_type"`
}

type RecoveryCodesRequest struct {
	UserId uint64 `json:"user_id" validate:"user_id"`
}

func (r SignUpRequest) Validate() error {

	_, err := govalidator.ValidateStruct(r)
	return rest.ValidateResponse(err)
}

func (r SignUpRequest) ValidateMail() error {
	_, err := govalidator.ValidateStruct(r)
	govalidator.IsEmail(r.Email)
	return rest.ValidateResponse(err)
}

func (r SignUpRequest) ValidatePhone() error {
	 err := isNumber(r.PhoneNumber+r.PhoneCountryCode)
	 if err!=nil{
	 	return err
	 }

	return err
}
func (r SignUpRequest) ValidateMfaType() error {

	if !((r.MfaType == "otp") || (r.MfaType == "phone") || (r.MfaType == "")) {

		return rest.ErrFieldResp{
			Meta: rest.ErrFieldRespMeta{
				ErrMessage: "Uncorrect multifactor authentification type",
			},
		}
	}
	return nil
}

// SignIn Request
type SignInRequest struct {
	Name     string `json:"name" validate:"name"`
	Password string `json:"password" validate:"password"`
}

func (r SignInRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return rest.ValidateResponse(err)
}

// SignOut Request
type SignOutRequest struct{}

func (r RecoveryCodesRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return rest.ValidateResponse(err)
}


func isNumber(s string) error {
	for _, r := range s {
		if (r < '0' || r > '9')  {
			return rest.ErrFieldResp{
				Meta: rest.ErrFieldRespMeta{
					ErrMessage: "Phone number has non-numeric symbol",
				},
			}
		}
	}
	return nil
}