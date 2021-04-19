package social_provider_status

type Status uint8

const (
	Disabled Status = iota
	Enabled
)

func (u Status) Value() uint8 {
	return uint8(u)
}

func New(status uint8) Status {
	return Status(status)
}
