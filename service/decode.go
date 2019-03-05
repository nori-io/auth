package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func DecodeSignUpRequest(parameters PluginParameters) func(_ context.Context, r *http.Request) (interface{}, error) {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var body SignUpRequest
		var isTypeValid bool

		if body.PhoneNumber==""{
			err := errors.New("Type '" + body.Type + "' is not valid ")
			return body, err
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return body, err
		}
		if err := body.Validate(); err != nil {
			return body, err
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
			err := errors.New("Type '" + body.Type + "' is not valid ")
			return body, err
		}

		if ((parameters.UserRegistrationPhoneNumberType)||(parameters.UserRegistrationEmailAddressType))!=true{
			err := errors.New(" All user's registration's types sets with 'false' value. Need to set 'true' value")
			return body, err
		}
		return body, nil
	}

}

func DecodeLogInRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body SignInRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return body, err
	}
	if err := body.Validate(); err != nil {
		return body, err
	}
	return body, nil
}

func DecodeLogOutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body SignOutRequest
	return body, nil
}
