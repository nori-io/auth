package error

import "net/http"

type ErrorHelper interface {
	Error(w http.ResponseWriter, r *http.Request, err error)
}
