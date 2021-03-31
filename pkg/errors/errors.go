package errors

import (
	"github.com/pkg/errors"
)

type Error struct {
	Err  error
	Key  string
	Type Type
}

func (e Error) Error() string {
	return e.Err.Error()
}

func New(key, msg string, t Type) Error {
	return Error{
		Err:  errors.WithStack(errors.New(msg)),
		Key:  key,
		Type: t,
	}
}

func NewInternal(err error) Error {
	return Error{
		Err:  errors.WithStack(err),
		Key:  "internal_server_error",
		Type: ErrInternal,
	}
}
