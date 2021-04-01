package cookie

import (
	"errors"
	"net/http"
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (h CookieHelper) GetSessionID(r *http.Request) (string, error) {
	cookie, err := r.Cookie("ssid")
	if err != nil {
		return "", err
	}
	if cookie == nil {
		// todo: error
		return "", errors.New("no cookie")
	}
	return cookie.Value, nil
}

func (h CookieHelper) SetSession(w http.ResponseWriter, session *entity.Session) {
	http.SetCookie(w, &http.Cookie{
		Name:   h.config.CookiesName(),
		Value:  string(session.SessionKey),
		Path:   h.config.CookiesPath(),
		Domain: h.config.CookiesDomain(),
		//@todo Expires
		Expires: time.Now().Add(time.Duration(h.config.CookiesExpires()) * time.Second),
		// MaxAge:   h.config.CookiesMaxAge(),
		Secure:   h.config.CookiesSecure(),
		HttpOnly: h.config.CookiesHttpOnly(),
		SameSite: http.SameSite(h.config.CookiesSameSite()),
	})
}

func (h CookieHelper) UnsetSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   h.config.CookiesName(),
		Value:  "",
		Path:   h.config.CookiesPath(),
		Domain: h.config.CookiesDomain(),
		//@todo Expires
		MaxAge:   -1,
		Secure:   h.config.CookiesSecure(),
		HttpOnly: h.config.CookiesHttpOnly(),
		SameSite: http.SameSite(h.config.CookiesSameSite()),
	})
}
