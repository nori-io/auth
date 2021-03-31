package http

import (
	"net/http"

	"github.com/nori-plugins/authentication/pkg/errors"
)

func Error(w http.ResponseWriter, r *http.Request, err error) {

	// h.logger.Error("%s", err)

	switch e := err.(type) {
	case errors.Error:
		switch e.Type {
		case errors.ErrInternal:
			//
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
