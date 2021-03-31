package errors

import "github.com/nori-plugins/authentication/pkg/errors"

var (
	UserNotFound  = errors.New("authentication.user_not_found", "user not found", errors.ErrNotFound)
	DuplicateUser = errors.New("authentication.duplicate_user", "duplicate user", errors.ErrConflict)
)
