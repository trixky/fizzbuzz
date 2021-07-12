package routes

import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"crypto/md5"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/json_struct"
	"fizzbuzz.com/v1/models"
	"github.com/julienschmidt/httprouter"
)

func Register(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	pseudo := req.FormValue("pseudo")
	password := req.FormValue("password")

	// check if we have pseudo and password
	if len(pseudo) < 1 || len(password) < 1 {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "pseudo or password missing/bad formatted"})
		return
	}

	// check if pseudo is available
	var pseudo_count int64

	if database.Postgres.Table("api_users").Where("pseudo = ?", pseudo).Count(&pseudo_count).Error != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	}

	if pseudo_count > 0 {
		res.WriteHeader(http.StatusConflict)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "a user with this pseudo already exists"})
		return
	}

	// hash the password
	password_hasheur := md5.New()
	password_hasheur.Write([]byte(password))
	hashed_password := password_hasheur.Sum(nil)

	// create the new user in postgres
	if database.Postgres.Create(&models.Api_users{Pseudo: pseudo, Password: hex.EncodeToString(hashed_password)}).Error != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	}
}
