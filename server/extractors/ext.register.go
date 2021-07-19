package extractors

import (
	"encoding/json"
	"errors"
	"net/http"

	"fizzbuzz.com/v1/tools"
)

type Register struct {
	Pseudo   string `json:"pseudo"`
	Password string `json:"password"`
}

// Extracts extracts parameters from the request
func (r *Register) Extracts(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&r); err != nil {
		return errors.New("invalid request body")
	}

	if len(r.Pseudo) < 3 || len(r.Pseudo) > 20 {
		return errors.New("pseudo must be between 3 and 20 characters")
	}

	if err := tools.Password_is_valid_v1(r.Password); err != nil {
		return err
	}

	return nil
}
