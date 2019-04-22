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
		var isTypeValid bool

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

		var errorResponse rest.ErrFieldResp

		if isTypeValid == false {

			errorResponse.AddError("type",0,"User type isn't valid")
		}

		if ((parameters.UserRegistrationPhoneNumberType) || (parameters.UserRegistrationEmailAddressType)) != true {



			}

		if (parameters.UserRegistrationEmailAddressType == true) && (parameters.UserRegistrationPhoneNumberType == false) {
			if body.ValidateMail()!=nil{

				errorResponse.AddError("email",0,
					"Mail address' format is uncorrect")
			}

		}

		if (parameters.UserRegistrationEmailAddressType == false) && (parameters.UserRegistrationPhoneNumberType == true) {
			if body.ValidatePhone()!=nil{
				errorResponse.AddError("phone_country_code",0,
					"Country code's format is uncorrect")

				errorResponse.AddError("phone_number",0,
					"Phone number's format is uncorrect ")

			}
		}

		if (parameters.UserRegistrationEmailAddressType == true) && (parameters.UserRegistrationPhoneNumberType == true) {
			if (body.ValidateMail()==nil)&&(body.ValidatePhone()==nil){

				errorResponse.AddError("email",0,
					"Mail address' format is uncorrect")
			}
			if body.ValidatePhone()!=nil{
				errorResponse.AddError("phone_country_code",0,
					"Country code's format is uncorrect")

				errorResponse.AddError("phone_number",0,
					"Phone number's format is uncorrect ")



			}

		}

		if parameters.UserMfaTypeParameter == "" {
			body.Validate()
		}

       fmt.Println("errorResponse.HasErrors",errorResponse.HasErrors())

		if errorResponse.HasErrors()==true {
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
