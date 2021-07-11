package middleware

import (
	"encoding/json"
	"net/http"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/json_struct"
	uuid "github.com/satori/go.uuid"
)

func Middleware_token(res http.ResponseWriter, req *http.Request) (string, bool) {
	// get token from cookie

	cookie, err := req.Cookie("session")

	// check we have the token
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "'session' cookie is corrupted or missing"})
		return "", false
	}

	token := cookie.Value
	// check token format (is uuid ?)
	_, err = uuid.FromString(token)
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "'session' cookie is corrupted"})
		return "", false
	}

	// check token validity from redis
	result := database.Redis.Get("token>" + token)
	if result.Err() != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "your token from 'session' cookie is not valid"})
		return "", false
	}

	return result.Val(), true
}
