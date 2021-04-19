package errors

type Type uint8

const (
	ErrValidation    Type = iota + 1 // 400
	ErrUnauthorized                  // 401
	ErrForbidden                     // 403
	ErrNotFound                      // 404
	ErrAlreadyExists                 //  409
	ErrInternal                      // 500
	// todo: add more types
)
