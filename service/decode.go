package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cheebo/gorest"
)

func DecodeSignUpRequest(parameters PluginParameters) func(_ context.Context, r *http.Request) (interface{}, error) {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var body SignUpRequest
		var errorResponse rest.ErrFieldResp

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return body, rest.ErrFieldResp{
				Meta: rest.ErrMeta{
					ErrMessage: "Error of decoding",
				},
			}
		}

		if err := body.Validate(); err != nil {
			return body, rest.ErrFieldResp{
				Meta: rest.ErrMeta{
					ErrMessage: "Error of body.Validate()",
				},
			}
		}
		if !userTypeValidate(parameters.UserTypeParameter, parameters.UserTypeDefaultParameter, body.Type) {
			errorResponse.AddError("type", 0, "User type isn't valid")
		}

		if (parameters.UserRegistrationByEmailAddress) && (!parameters.UserRegistrationByPhoneNumber) {
			if !body.ValidateMail() {

				errorResponse.AddError("email", 0,
					"Mail address' format is Incorrect")
			}

		}

		if (!parameters.UserRegistrationByEmailAddress) && (parameters.UserRegistrationByPhoneNumber) {
			errPhoneCountryCode, errPhoneNumber := body.ValidatePhone()

			if errPhoneCountryCode != nil {
				errorResponse.AddError("phone_country_code", 0,
					"Country code's format is Incorrect")

			}

			if errPhoneNumber != nil {
				errorResponse.AddError("phone_number", 0,
					"Phone number's format is Incorrect ")
			}

		}

		if (parameters.UserRegistrationByEmailAddress) && (parameters.UserRegistrationByPhoneNumber) {
			errPhoneCountryCode, errPhoneNumber := body.ValidatePhone()
			if body.Email != "" {
				if !body.ValidateMail() {
					errorResponse.AddError("email", 0,
						"Mail address' format is Incorrect")
				}
			}
			if len(body.PhoneCountryCode+body.PhoneNumber) != 0 {
				if errPhoneCountryCode != nil {
					errorResponse.AddError("phone_country_code", 0,
						"Country code's format is Incorrect")

				}
			}

			if errPhoneNumber != nil {
				errorResponse.AddError("phone_number", 0,
					"Phone number's format is Incorrect ")
			}

		}

		if len(body.Password) == 0 {
			errorResponse.AddError("password", 0,
				"Password is empty")
		}

		if !((parameters.UserMfaTypeParameter == "") || (parameters.UserMfaTypeParameter == MfaTypeSMS) || (parameters.UserMfaTypeParameter == MfaTypeOTP)) {
			errorResponse.AddError("mfa_type", 0,
				"Incorrect mfa type")
		}

		if len(body.Meta) == 0 {
			errorResponse.AddError("meta", 0,
				"Meta is absent")
		}

		if errorResponse.HasErrors() {
			return body, errorResponse
		}

		return body, nil
	}

}

func DecodeSignInRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body SignInRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return body, err
	}
	if err := body.Validate(); err != nil {
		return body, err
	}
	return body, nil
}

func DecodeSignOutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body SignOutRequest

	return body, nil
}

func DecodeRecoveryCodes() func(_ context.Context, r *http.Request) (interface{}, error) {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var body RecoveryCodesRequest
		var errorText string
		var errCommon error

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return body, rest.ErrFieldResp{
				Meta: rest.ErrMeta{
					ErrMessage: "Error of decoding",
				},
			}
		}
		if err := body.Validate(); err != nil {
			return body, rest.ErrFieldResp{
				Meta: rest.ErrMeta{
					ErrMessage: "Error of body.Validate()",
				},
			}
		}

		if len(errorText) != 0 {
			return body, rest.ErrFieldResp{
				Meta: rest.ErrMeta{
					ErrMessage: errCommon.Error(),
				},
			}
		}

		return body, nil
	}

}

func userTypeValidate(userType []interface{}, userTypeDefault, bodyType string) bool {
	var isTypeValid bool

	typesSlice := userType
	s := make([]string, len(typesSlice))
	for i, value := range typesSlice {
		s[i] = fmt.Sprint(value)
	}

	for _, value := range s {
		if value == bodyType {
			isTypeValid = true
		}
	}

	if userTypeDefault == bodyType {
		isTypeValid = true
	}

	if bodyType == "" {
		bodyType = userTypeDefault
		isTypeValid = true
	}
	if !isTypeValid {

		return false
	}
	return true
}
