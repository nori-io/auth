package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cheebo/gorest"
)

func DecodeSignUpRequest(parameters PluginParameters) func(_ context.Context, r *http.Request) (interface{}, error) {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var body SignUpRequest
		var isTypeValid bool
		var errorText string
		var errCommon error

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return body, rest.ErrFieldResp{
				Meta: rest.ErrFieldRespMeta{
					ErrMessage: "Error of decoding",
				},
			}
		}

		if err := body.Validate(); err != nil {
			return body, rest.ErrFieldResp{
				Meta: rest.ErrFieldRespMeta{
					ErrMessage: "Error of body.Validate()",
				},
			}
		}

		typesSlice := parameters.UserTypeParameter
		s := make([]string, len(typesSlice))
		for i, value := range typesSlice {
			s[i] = fmt.Sprint(value)
		}

		for _, value := range s {
			if value == body.Type {
				isTypeValid = true
			}
		}

		if parameters.UserTypeDefaultParameter == body.Type {
			isTypeValid = true
		}

		if body.Type == "" {
			body.Type = parameters.UserTypeDefaultParameter
			isTypeValid = true
		}
		if isTypeValid == false {
			errorText = errorText + "Type '" + body.Type + "' is not valid \n"
			errCommon = errors.New(errorText)
		}

		if ((parameters.UserRegistrationPhoneNumberType) || (parameters.UserRegistrationEmailAddressType)) != true {
			errorText = errorText + " All user's registration's types sets with 'false' value. Need to set 'true' value \n "
			errCommon = errors.New(errorText)
		}

		if (parameters.UserRegistrationEmailAddressType == true) && (parameters.UserRegistrationPhoneNumberType == false) {
			body.ValidateOnlyByMail()

		}
		if (parameters.UserRegistrationEmailAddressType == true) && (parameters.UserRegistrationPhoneNumberType == false) {
			body.ValidateOnlyByPhone()
		}

		if (parameters.UserRegistrationEmailAddressType == true) && (parameters.UserRegistrationPhoneNumberType == true) {
			body.Validate()
		}

		if errorText != "" {
			return body, rest.ErrFieldResp{
				Meta: rest.ErrFieldRespMeta{
					ErrMessage: errCommon.Error(),
				},
			}
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
				Meta: rest.ErrFieldRespMeta{
					ErrMessage: "Error of decoding",
				},
			}
		}
		if err := body.Validate(); err != nil {
			return body, rest.ErrFieldResp{
				Meta: rest.ErrFieldRespMeta{
					ErrMessage: "Error of body.Validate()",
				},
			}
		}

		if errorText != "" {
			return body, rest.ErrFieldResp{
				Meta: rest.ErrFieldRespMeta{
					ErrMessage: errCommon.Error(),
				},
			}
		}

		return body, nil
	}

}
