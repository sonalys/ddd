package interfaces

import (
	"encoding/json"
	"net/http"
)

// RequestBuilder determines a struct that can be decoded and validated from a request.
type RequestBuilder interface {
	Decode(*http.Request) error
	Valid() error
}

func writeJSON(w http.ResponseWriter, data interface{}, code int) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(buf)
		w.WriteHeader(code)
	}
	return nil
}

func parseRequest(r *http.Request, d RequestBuilder) error {
	if err := d.Decode(r); err != nil {
		return err
	}

	if err := d.Valid(); err != nil {
		return err
	}

	return nil
}
