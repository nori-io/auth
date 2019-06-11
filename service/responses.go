package service

// SignUpResponse
type SignUpResponse struct {
	Email            string
	PhoneCountryCode string
	PhoneNumber      string
	HttpStatusCode   int
	Err              error
}

func (d *SignUpResponse) Error() error {
	return d.Err
}

func (d *SignUpResponse) StatusCode() int {
	return d.HttpStatusCode
}

// SignInResponse
type SignInResponse struct {
	Id             uint64
	Token          string
	User           UserResponse
	MFA            string
	HttpStatusCode int
	Err            error
}

type UserResponse struct {
	UserName string
}

func (d *SignInResponse) Error() error {
	return d.Err
}

func (d *SignInResponse) StatusCode() int {
	return d.HttpStatusCode
}

// SignOut Response
type SignOutResponse struct {
	HttpStatusCode int
	Err            error
}

func (d *SignOutResponse) Error() error {
	return d.Err
}

func (d *SignOutResponse) StatusCode() int {
	return d.HttpStatusCode
}

type RecoveryCodesResponse struct {
	Codes          []string
	HttpStatusCode int
	Err            error
}

func (d *RecoveryCodesResponse) Error() error {
	return d.Err
}

func (d *RecoveryCodesResponse) StatusCode() int {
	return d.HttpStatusCode
}

type SignInSocialResponse struct {
	Id             uint64
	Token          string
	User           UserResponse
	MFA            string
	HttpStatusCode int
	Err            error
}

func (d *SignInSocialResponse) Error() error {
	return d.Err
}

func (d *SignInSocialResponse) StatusCode() int {
	return d.HttpStatusCode
}

type SignOutSocialResponse struct {
	HttpStatusCode int
	Err            error
}

func (d *SignOutSocialResponse) Error() error {
	return d.Err
}

func (d *SignOutSocialResponse) StatusCode() int {
	return d.HttpStatusCode
}
