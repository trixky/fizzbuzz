package routes

import (
	"encoding/json"
	"net/http"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/json_struct"
	"fizzbuzz.com/v1/models"
	"github.com/julienschmidt/httprouter"
)

type Data_stat struct {
	Int1           int    `json:"int1"`
	Int2           int    `json:"int12"`
	Limit          int    `json:"limit"`
	Str1           string `json:"str1"`
	Str2           string `json:"str2"`
	Request_number int    `json:"request_number"`
}

type Data_stats struct {
	Requests []Data_stat `json:"request"`
}

func Stats(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	reqsss := []models.Fizzbuzz_request_statistics{}

	result := database.Postgres.Order("request_number DESC").Limit(10).Find(&reqsss)
	if result.Error != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	}

	data_stats := Data_stats{}
	for _, request := range reqsss {
		data_stats.Requests = append(data_stats.Requests, Data_stat{request.Int1, request.Int2, request.Limit, request.Str1, request.Str2, request.Request_number})
	}

	json.NewEncoder(res).Encode(data_stats)
}
