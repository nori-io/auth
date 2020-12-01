package service

type AuthenticationService interface {
	SignUp(ctx context.Context, data SignUpData) (*entity.User, error)
	SignIn(ctx context.Context, data SignInData) (*entity.Session, error)
	SignOut(ctx context.Context, data SignOutData) error
}

type SignUpData struct {
	Email    string
	Password string
}

func (d SignUpData) Validate() error {
	return nil
}
