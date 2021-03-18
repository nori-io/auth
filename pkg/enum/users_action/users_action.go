package users_action

type Action uint8

const (
	SignUp Action = iota
	SignIn
	SignOut
)

func (u Action) Value() uint8 {
	return uint8(u)
}

func New(action uint8) Action {
	return Action(action)
}
