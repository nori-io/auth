package users_type

type UserType uint8

const (
	Admin UserType = iota
	User
	Guest
)

func (u UserType) Value() uint8 {
	return uint8(u)
}

func New(userType uint8) UserType {
	return UserType(userType)
}
