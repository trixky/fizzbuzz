package routes

import (
	"encoding/json"
	"net/http"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/json_struct"
	"fizzbuzz.com/v1/models"
	"github.com/julienschmidt/httprouter"
)

// Data_stat contains all the interesting information about the most popular fizzbuzz requests.
type Data_stat struct {
	Int1           int    `json:"int1"`
	Int2           int    `json:"int2"`
	Limit          int    `json:"limit"`
	Str1           string `json:"str1"`
	Str2           string `json:"str2"`
	Request_number int    `json:"request_number"`
}

// Data_stats contains the list of the most popular fizzbuzz requests.
type Data_stats struct {
	Requests []Data_stat `json:"request"`
}

// Stats handles the [GET /stats] request.
func Stats(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	results := []models.Fizzbuzz_request_statistics{}

	if err := database.Postgres.Order("request_number DESC").Limit(10).Find(&results).Error; err != nil {
		// If postgres failed (to deliver the most popular fizzbuzz requests).
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	}

	data_stats := Data_stats{}

	for _, request := range results {
		// encapsulates the most popular fizzbuzz requests.
		data_stats.Requests = append(data_stats.Requests, Data_stat{request.Int1, request.Int2, request.Limit, request.Str1, request.Str2, request.Request_number})
	}

	json.NewEncoder(res).Encode(data_stats)
}
