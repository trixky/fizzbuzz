package extractors

import (
	"encoding/json"
	"errors"
	"net/http"

	"fizzbuzz.com/v1/tools"
)

type Extracted_login struct {
	Pseudo   string `json:"pseudo"`
	Password string `json:"password"`
}

type Extracted_register struct {
	Pseudo   string `json:"pseudo"`
	Password string `json:"password"`
}

type Extracted_block struct {
	Pseudo       string `json:"pseudo"`
	Block_status string `json:"block_status"` // "true" or "false"
}

func Extract_login(req *http.Request) (Extracted_login, error) {
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

func Extract_block(req *http.Request) (Extracted_block, error) {
	extracted_block := Extracted_block{}

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&extracted_block); err != nil {
		return extracted_block, errors.New("invalid request body")
	}

	if len(extracted_block.Pseudo) < 1 {
		return extracted_block, errors.New("pseudo is missing")
	}

	if len(extracted_block.Block_status) < 1 {
		return extracted_block, errors.New("block_status is missing")
	}

	if extracted_block.Block_status != "true" && extracted_block.Block_status != "false" {
		return extracted_block, errors.New("block_status must have the value [true/false]")
	}

	return extracted_block, nil
}
