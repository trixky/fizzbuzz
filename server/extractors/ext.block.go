// Package extractors allows to extract the parameters of requests
package extractors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Block struct {
	Pseudo       string `json:"pseudo"`
	Block_status string `json:"block_status"` // "true" or "false"
}

// Extracts extracts parameters from the request
func (b *Block) Extracts(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&b); err != nil {
		return errors.New("invalid request body")
	}

	if len(b.Pseudo) < 1 {
		return errors.New("pseudo is missing")
	}

	if len(b.Block_status) < 1 {
		return errors.New("block_status is missing")
	}

	if b.Block_status != "true" && b.Block_status != "false" {
		return errors.New("block_status must have the value [true/false]")
	}

	return nil
}
