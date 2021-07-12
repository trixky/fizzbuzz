package routes

import (
	"encoding/json"
	"net/http"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/json_struct"
	"fizzbuzz.com/v1/models"
	"fizzbuzz.com/v1/tools"
	"github.com/julienschmidt/httprouter"
)

// Register handles the [POST /register] request.
func Register(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	pseudo := req.FormValue("pseudo")
	password := req.FormValue("password")

	if len(pseudo) < 1 || len(password) < 1 {
		// If pseudo or password missing/bad formatted.
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "pseudo or password missing/bad formatted"})
		return
	}

	var pseudo_count int64

	if database.Postgres.Table("api_users").Where("pseudo = ?", pseudo).Count(&pseudo_count).Error != nil {
		// If postgres failed to search a user with this pseudo.
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	}

	if pseudo_count > 0 {
		// If a user with this pseudo already exists.
		res.WriteHeader(http.StatusConflict)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "a user with this pseudo already exists"})
		return
	}

	hashed_password_str := tools.Hash_password(password) // hash the password

	if database.Postgres.Create(&models.Api_users{Pseudo: pseudo, Password: hashed_password_str}).Error != nil {
		// If postgres failed to create the new user.
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	}
}
