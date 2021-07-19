package extractors

import (
	"errors"
	"net/http"
	"strconv"
)

type Fizzbuzz struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// Extracts extracts parameters from the request
func (f *Fizzbuzz) Extracts(req *http.Request) error {
	f.Int1, _ = strconv.Atoi(req.URL.Query().Get("int1"))
	f.Int2, _ = strconv.Atoi(req.URL.Query().Get("int2"))
	f.Limit, _ = strconv.Atoi(req.URL.Query().Get("limit"))
	f.Str1 = req.URL.Query().Get("str1")
	f.Str2 = req.URL.Query().Get("str2")

	if f.Int1 < 1 {
		return errors.New("int1 is missing or not strictly positive or bad formated")
	}

	if f.Int2 < 1 {
		return errors.New("int2 is missing or not strictly positive or bad formated")
	}

	if f.Limit < 1 {
		return errors.New("limit is missing or not strictly positive or bad formated")
	}

	if f.Int1 == f.Int2 {
		return errors.New("int1 can't be equal to int2")
	}

	if f.Int1 > f.Limit {
		return errors.New("int1 cannot be greater than limit")
	}

	if f.Int2 > f.Limit {
		return errors.New("int2 cannot be greater than limit")
	}

	if len(f.Str1) < 1 {
		return errors.New("str1 is missing")
	}

	if len(f.Str1) > 30 {
		return errors.New("str1 cannot be more than 30 characters")
	}

	if len(f.Str2) < 1 {
		return errors.New("str2 is missing")
	}

	if len(f.Str1) > 30 {
		return errors.New("str1 cannot be more than 30 characters")
	}

	return nil
}
