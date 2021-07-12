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

// Block handles the [PATCH /block] request.
func Block(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	pseudo, _ := req.Context().Value(middlewares.Key_middleware_infos).(*middlewares.Middleware_infos).Get("pseudo") // ignore err

	subject_pseudo := req.FormValue("pseudo")     // get input pseudo
	block_status := req.FormValue("block_status") // get input block_status

	if len(subject_pseudo) < 1 || (block_status != "true" && block_status != "false") {
		// If pseudo and/or block_status input(s) missing/bad formatted.
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "pseudo and/or block_status missing/bad formatted"})
		return
	}

	api_user := models.Api_users{}

	if err := database.Postgres.Table("api_users").Where(map[string]interface{}{"pseudo": pseudo, "admin": true}).First(&api_user).Error; err != nil {
		// If postgres could not find the user.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If user don't have the required privileges.
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "you don't have the required privileges"})
		} else {
			// Else postgres failed.
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		}
		return
	}

	if result := database.Postgres.Table("api_users").Where("pseudo = ?", subject_pseudo).Update("blocked", block_status); result.Error != nil {
		// If postgres failed to update the 'blocked' status.
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	} else if result.RowsAffected > 0 {
		if block_status == "true" {
			// If the new 'blocked' status is 'true'.
			tokens_key := "tokens>" + subject_pseudo
			if result := database.Redis.Keys(tokens_key); result.Err() != nil {
				// If redis failed to get token keys.
				res.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
				return
			} else if len(result.Val()) > 0 {
				result := database.Redis.Get(tokens_key)
				if result.Err() != nil {
					// If redis failed to get token values.
					res.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
					return
				}
				tokens := strings.Split(result.Val(), "|")[1:]
				for _, token := range tokens {
					database.Redis.Del("token>" + token) // delete each token key
				} // ignore err
				database.Redis.Del(tokens_key) // delete token keys
			}
		}
	} else {
		// Else not user finded with this pseudo.
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "bad pseudo"})
	}
}
