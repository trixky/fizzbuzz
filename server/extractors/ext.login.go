package extractors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Login struct {
	Pseudo   string `json:"pseudo"`
	Password string `json:"password"`
}

// Extracts extracts parameters from the request
func (l *Login) Extracts(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&l); err != nil {
		return errors.New("invalid request body")
	}

	if len(l.Pseudo) < 1 {
		return errors.New("pseudo is missing")
	}

	if len(l.Password) < 1 {
		return errors.New("password is missing")
	}

	return nil
}
