package users_action

type Action uint8

const (
	EnableMfa Action = iota
	DisableMfa
	SignUp
	LogIn
	LogInMfa
	LogOut
)

func (u Action) Value() uint8 {
	return uint8(u)
}

func New(action uint8) Action {
	return Action(action)
}
