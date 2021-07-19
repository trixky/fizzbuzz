package logic

import "fizzbuzz.com/v1/models"

// Fizzbuzz_request_statistic contains all the interesting information about the most popular fizzbuzz requests.
type Fizzbuzz_request_statistic struct {
	Int1           int    `json:"int1"`
	Int2           int    `json:"int2"`
	Limit          int    `json:"limit"`
	Str1           string `json:"str1"`
	Str2           string `json:"str2"`
	Request_number int    `json:"request_number"`
}

// Fizzbuzz_request_statistics contains the list of the most popular fizzbuzz requests.
type Fizzbuzz_request_statistics struct {
	Requests []Fizzbuzz_request_statistic `json:"requests"`
}

// Fizzbuzz_request_statistics_generator generates a list of statistics about fizzbuzz requests.
func Fizzbuzz_request_statistics_generator(results []models.Fizzbuzz_request_statistics) (fizzbuzz_request_statistics Fizzbuzz_request_statistics) {
	for _, request := range results {
		// encapsulates the most popular fizzbuzz requests.
		fizzbuzz_request_statistics.Requests = append(fizzbuzz_request_statistics.Requests, Fizzbuzz_request_statistic{request.Int1, request.Int2, request.Limit, request.Str1, request.Str2, request.Request_number})
	}

	return
}
