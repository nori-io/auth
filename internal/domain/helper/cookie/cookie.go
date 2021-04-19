package cookie

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type CookieHelper interface {
	GetSessionID(r *http.Request) (string, error)
	SetSession(w http.ResponseWriter, session *entity.Session)
	UnsetSession(w http.ResponseWriter)
}
