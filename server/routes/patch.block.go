package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/json_struct"
	"fizzbuzz.com/v1/middlewares"
	"fizzbuzz.com/v1/models"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Block(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	// ignore err
	pseudo, _ := req.Context().Value(middlewares.Key_middleware_infos).(*middlewares.Middleware_infos).Get_pseudo("pseudo")

	subject_pseudo := req.FormValue("pseudo")
	block_status := req.FormValue("block_status")

	if len(subject_pseudo) < 1 || (block_status != "true" && block_status != "false") {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "pseudo and/or block_status missing/bad formatted"})
		return
	}

	api_user := models.Api_users{}

	if result := database.Postgres.Table("api_users").Where(map[string]interface{}{"pseudo": pseudo, "admin": true}).First(&api_user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "you do not have the required privileges"})
		} else {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		}
		return
	}

	if result := database.Postgres.Table("api_users").Where("pseudo = ?", subject_pseudo).Update("blocked", block_status); result.Error != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	} else if result.RowsAffected > 0 {
		if block_status == "true" {
			tokens_key := "tokens>" + subject_pseudo
			if result := database.Redis.Keys(tokens_key); result.Err() != nil {
				res.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
				return
			} else if len(result.Val()) > 0 {
				result := database.Redis.Get(tokens_key)
				if result.Err() != nil {
					res.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
					return
				}
				tokens := strings.Split(result.Val(), "|")[1:]
				for _, token := range tokens {
					database.Redis.Del("token>" + token)
				}
				database.Redis.Del(tokens_key)
			}
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "bad pseudo"})
	}
}
