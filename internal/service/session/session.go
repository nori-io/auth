package session

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"
	"github.com/nori-plugins/authentication/pkg/errors"
)

func (srv SessionService) IsSessionExist(ctx context.Context, sessionKey string) (bool, error) {
	session, err := srv.sessionRepository.FindBySessionKey(ctx, sessionKey)
	if err != nil {
		return false, errors.NewInternal(err)
	}
	if session != nil && session.Status == session_status.Active {
		return true, errors.New("user_session_exists", "users already sign in, sign up isn't possible", errors.ErrConflict)
	}
	return false, nil
}
