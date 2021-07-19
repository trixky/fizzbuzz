// Package controllers gives all the handler functions corresponding to their routes and methods.
package controllers

import (
	"encoding/json"
	"net/http"

	"fizzbuzz.com/v1/extractors"
	"fizzbuzz.com/v1/logic"
	"fizzbuzz.com/v1/repositories"
	"fizzbuzz.com/v1/tools"
	"github.com/julienschmidt/httprouter"
)

// Fizzbuzz handles the [GET /fizzbuzz] request.
func Fizzbuzz(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	// #extraction
	extractor := extractors.Fizzbuzz{}
	if err := extractor.Extracts(req); err != nil {
		// If extraction failed.
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(tools.Data_error{Error: err.Error()})
		return
	}

	// #repository [create or increment fizzbuzz_request_statistics]
	if _, err := repositories.Create_or_increment_fizzbuzz_request_statistics(extractor.Int1, extractor.Int2, extractor.Limit, extractor.Str1, extractor.Str2); err != nil {
		// If repository failed.
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(tools.Data_error{Error: "internal server error"})
		return
	}

	// #logic [generates fizzbuzz]
	data_fizzbuzz := logic.Fizzbuzz_generator(&extractor)

	// $ success response $
	json.NewEncoder(res).Encode(data_fizzbuzz)
}

// Stats handles the [GET /stats] request.
func Stats(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")
	// #repository [get all fizzbuzz_request_statistics]
	results, err := repositories.Getall_fizzbuzz_request_statistics()
	if err != nil {
		// If repository failed.
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(tools.Data_error{Error: "internal server error"})
		return
	}

	// #logic [generates fizzbuzz]
	fizzbuzz_request_statistics := logic.Fizzbuzz_request_statistics_generator(results)

	// $ success response $
	json.NewEncoder(res).Encode(fizzbuzz_request_statistics)
}
