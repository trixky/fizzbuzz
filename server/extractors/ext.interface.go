package extractors

import (
	"net/http"
)

type Extractor interface {
	Extracts(req *http.Request) error
}
