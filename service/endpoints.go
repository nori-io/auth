package service

import (
	"context"
	"github.com/nori-io/nori-common/endpoint"
)

func MakeSignUpEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(SignUpRequest)
		resp := s.SignUp(ctx, req)
		return *resp, resp.Error()
	}

	return nil
}

func MakeSignInEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {

		req := r.(SignInRequest)
		resp := s.SignIn(ctx, req)
		return *resp, resp.Error()
	}
	return nil
}

func MakeSignOutEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(SignOutRequest)
		resp := s.SignOut(ctx, req)
		return *resp, resp.Error()
	}
	return nil
}
