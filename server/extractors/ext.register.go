package extractors

import (
	"encoding/json"
	"errors"
	"net/http"

	"fizzbuzz.com/v1/tools"
)

type Extracted_register struct {
	Pseudo   string `json:"pseudo"`
	Password string `json:"password"`
}

func Extract_register(req *http.Request) (extracted_register Extracted_register, err error) {
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&extracted_register); err != nil {
		return extracted_register, errors.New("invalid request body")
	}

	if len(extracted_register.Pseudo) < 3 || len(extracted_register.Pseudo) > 20 {
		return extracted_register, errors.New("pseudo must be between 3 and 20 characters")
	}

	if len(extracted_register.Password) < 8 || len(extracted_register.Password) > 30 {
		return extracted_register, errors.New("password must be between 8 and 30 characters")
	}

	if err := tools.Password_is_valid_v1(extracted_register.Password); err != nil {
		return extracted_register, err
	}

	return extracted_register, nil
}
