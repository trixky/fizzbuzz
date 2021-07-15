package extractors

import (
	"errors"
	"net/http"
	"strconv"
)

type Extracted_fizzbuzz struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// Extracts_block extracts parameters frome the /fizzbuzz endpoint
func Extracts_fizzbuzz(req *http.Request) (Extracted_fizzbuzz, error) {
	extracted_fizzbuzz := Extracted_fizzbuzz{}

	extracted_fizzbuzz.Int1, _ = strconv.Atoi(req.URL.Query().Get("int1"))
	extracted_fizzbuzz.Int2, _ = strconv.Atoi(req.URL.Query().Get("int2"))
	extracted_fizzbuzz.Limit, _ = strconv.Atoi(req.URL.Query().Get("limit"))
	extracted_fizzbuzz.Str1 = req.URL.Query().Get("str1")
	extracted_fizzbuzz.Str2 = req.URL.Query().Get("str2")

	if extracted_fizzbuzz.Int1 < 1 {
		return extracted_fizzbuzz, errors.New("int1 is missing or not strictly positive or bad formated")
	}

	if extracted_fizzbuzz.Int2 < 1 {
		return extracted_fizzbuzz, errors.New("int2 is missing or not strictly positive or bad formated")
	}

	if extracted_fizzbuzz.Limit < 1 {
		return extracted_fizzbuzz, errors.New("limit is missing or not strictly positive or bad formated")
	}

	if extracted_fizzbuzz.Int1 == extracted_fizzbuzz.Int2 {
		return extracted_fizzbuzz, errors.New("int1 can't be equal to int2")
	}

	if extracted_fizzbuzz.Int1 > extracted_fizzbuzz.Limit {
		return extracted_fizzbuzz, errors.New("int1 cannot be greater than limit")
	}

	if extracted_fizzbuzz.Int2 > extracted_fizzbuzz.Limit {
		return extracted_fizzbuzz, errors.New("int2 cannot be greater than limit")
	}

	if len(extracted_fizzbuzz.Str1) < 1 {
		return extracted_fizzbuzz, errors.New("str1 is missing")
	}

	if len(extracted_fizzbuzz.Str1) > 30 {
		return extracted_fizzbuzz, errors.New("str1 cannot be more than 30 characters")
	}

	if len(extracted_fizzbuzz.Str2) < 1 {
		return extracted_fizzbuzz, errors.New("str2 is missing")
	}

	if len(extracted_fizzbuzz.Str1) > 30 {
		return extracted_fizzbuzz, errors.New("str1 cannot be more than 30 characters")
	}

	return extracted_fizzbuzz, nil
}
