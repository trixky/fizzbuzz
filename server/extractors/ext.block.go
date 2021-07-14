package extractors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Extracted_block struct {
	Pseudo       string `json:"pseudo"`
	Block_status string `json:"block_status"` // "true" or "false"
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
