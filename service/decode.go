package service

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeSignUpRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body SignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return body, err
	}
	if err := body.Validate(); err != nil {
		return body, err
	}
	return body, nil
}

func DecodeLogInRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return body, err
	}
	if err := body.Validate(); err != nil {
		return body, err
	}
	return body, nil
}

func DecodeLogOutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body LogoutRequest
	return body, nil
}

