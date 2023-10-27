package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	e "github.com/idekpas/xm-company/pkg/errors"
	"io"
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
		status = http.StatusInternalServerError
		body = ErrorBody{ErrorMessage: v.Error()}
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

func (s service) readRequestBody(r *http.Request) ([]byte, error) {
	// Read the content
	var bodyBytes []byte
	var err error
	if r.Body != nil {
		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			err := errors.New("could not read request body")
			return nil, err
		}
	}
	return bodyBytes, nil
}

func (s service) restoreRequestBody(r *http.Request, bodyBytes []byte) {
	// Restore the io.ReadCloser to its original state
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
}
