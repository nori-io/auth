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

		var errorResponse *rest.ErrFieldResp
		var errorObjects []rest.ErrFieldObject

		if isTypeValid == false {
			errorObject:=rest.ErrFieldObject{Message:"User type isn't valid"}
			errorObjects=append(errorObjects,errorObject)
			errorType:=rest.ErrField{Field:"type",Errs:errorObjects}
			fmt.Println(errorType)
		//	rest.ErrFieldResp.AddError(errorType,"",0,"")
		}

		if ((parameters.UserRegistrationPhoneNumberType) || (parameters.UserRegistrationEmailAddressType)) != true {
			errorObject:=rest.ErrFieldObject{Message:"All user's registration's types sets with 'false' value. Need to set 'true' value"}
			errorObjects=append(errorObjects,errorObject)
			errorRegistrationType:=rest.ErrField{Field:"parameters.UserRegistrationPhoneNumberType or parameters.UserRegistrationEmailAddressType in config file",Errs:errorObjects}
			rest.ErrFieldResp.AddError(errorRegistrationType,"",0,"")


		}

		if (parameters.UserRegistrationEmailAddressType == true) && (parameters.UserRegistrationPhoneNumberType == false) {
			if body.ValidateMail()!=nil{
				errorObject:=rest.ErrFieldObject{Message:"Mail address' format is uncorrect"}
				errorObjects=append(errorObjects,errorObject)
				errorMailFormat:=rest.ErrField{Field:"email",Errs:errorObjects}
				rest.ErrFieldResp.AddError(errorMailFormat,"",0,"")


			}

		}
		if (parameters.UserRegistrationEmailAddressType == false) && (parameters.UserRegistrationPhoneNumberType == true) {
			if body.ValidatePhone()!=nil{
				errorObjectCountryCode:=rest.ErrFieldObject{Message:"Country code's format is uncorrect "}
				errorObjectPhoneNumber:=rest.ErrFieldObject{Message:"Phone number's format is uncorrect "}

				errorObjects=append(errorObjects,errorObjectCountryCode)
				errorObjects=append(errorObjects,errorObjectPhoneNumber)

				errorMailFormat:=rest.ErrField{Field:"phone_country_code and phone_number",Errs:errorObjects}
				rest.ErrFieldResp.AddError(errorMailFormat,"",0,"")
			}
		}
/*
		if (parameters.UserRegistrationEmailAddressType == true) && (parameters.UserRegistrationPhoneNumberType == true) {
			if (body.ValidateMail()==nil)&&(body.ValidatePhone()==nil){
				if body.ValidateMail()!=nil{
					errorText = errorText + "Mail address' format is uncorrect "
				}
			}
			if body.ValidatePhone()!=nil{
				rest.ErrFieldResp.AddError("phone_number",)
				errorText = errorText + "Phone number's format is uncorrect\n"
				fmt.Println(errorText)



			}

		}*/

		if parameters.UserMfaTypeParameter == "" {
			body.Validate()
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
