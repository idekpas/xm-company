package handlers

import (
	"encoding/json"
	e "github.com/idekpas/xm-company/pkg/errors"
	"net/http"
)

func (s service) decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s service) respond(w http.ResponseWriter, data interface{}, status int) {
	var body interface{}
	switch v := data.(type) {
	case nil:
	case e.ErrArgument:
		status = http.StatusBadRequest
		body = ErrorBody{ErrorMessage: v.Unwrap().Error()}
	case error:
		if http.StatusText(status) == "" {
			status = http.StatusInternalServerError
		} else {
			body = ErrorBody{ErrorMessage: v.Error()}
		}
	default:
		body = data
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		err := json.NewEncoder(w).Encode(body)
		if err != nil {
			http.Error(w, "Could not encode in json", http.StatusBadRequest)
			return
		}
	}
}
