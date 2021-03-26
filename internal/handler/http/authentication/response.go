package authentication

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SignInResponse struct {
	MfaType string
}

func JSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Context-Type", "application/json; charset=utf-8")
	w.Write(buf.Bytes())
}
