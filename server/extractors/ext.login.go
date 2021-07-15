package extractors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Extracted_login struct {
	Pseudo   string `json:"pseudo"`
	Password string `json:"password"`
}

// Extracts_block extracts parameters frome the /login endpoint
func Extracts_login(req *http.Request) (Extracted_login, error) {
	extracted_login := Extracted_login{}

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&extracted_login); err != nil {
		return extracted_login, errors.New("invalid request body")
	}

	if len(extracted_login.Pseudo) < 1 {
		return extracted_login, errors.New("pseudo is missing")
	}

	if len(extracted_login.Password) < 1 {
		return extracted_login, errors.New("password is missing")
	}

	return extracted_login, nil
}
