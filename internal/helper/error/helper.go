package error

import (
	"encoding/json"
	"net/http"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nori-plugins/authentication/pkg/errors"
)

func (e errorHelper) Error(w http.ResponseWriter, err error) {
	e.logger.Error("%s", err)

	switch e := err.(type) {
	case errors.Error:
		switch e.Type {
		case errors.ErrValidation:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.ErrUnauthorized:
			http.Error(w, err.Error(), http.StatusUnauthorized)
		case errors.ErrForbidden:
			http.Error(w, err.Error(), http.StatusForbidden)
		case errors.ErrNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case errors.ErrAlreadyExists:
			http.Error(w, err.Error(), http.StatusConflict)
		case errors.ErrInternal:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case v.Errors:
		data, err := json.Marshal(e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Error(w, string(data), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
