package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/json_struct"
	"fizzbuzz.com/v1/models"
	"fizzbuzz.com/v1/tools"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"

	"gorm.io/gorm"
)

// Login handles the [POST /login] request.
func Login(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	pseudo := req.FormValue("pseudo")
	password := req.FormValue("password")

	// check if we have pseudo and password
	if len(pseudo) < 1 || len(password) < 1 {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "pseudo or password missing/bad formatted"})
		return
	}

	hashed_password_str := tools.Hash_password(password) // hash the password

	api_user := models.Api_users{}

	if err := database.Postgres.Table("api_users").Where(map[string]interface{}{"pseudo": pseudo, "password": hashed_password_str}).First(&api_user).Error; err != nil {
		// If postgres could not find the user.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If postgres could not find the user with this pseudo or/and password.
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "bad pseudo or/and password"})
		} else {
			// If postgres failed.
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		}
		return
	}

	if api_user.Blocked {
		// If user is blocked.
		http.SetCookie(res, &http.Cookie{
			Name:   "session",
			MaxAge: -1,
		})
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "your account is blocked"})
		return
	}

	new_session_token := uuid.NewV4().String() // generates a new session token (uuid)
	if database.Redis.Append("tokens>"+pseudo, "|"+new_session_token).Err() != nil {
		// If redis failed to add the token key to token keys.
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	}
	if database.Redis.Set("token>"+new_session_token, pseudo, 0).Err() != nil {
		// If redis failed to add the token key.
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	}

	http.SetCookie(res, &http.Cookie{
		Name:  "session",
		Value: new_session_token,
	}) // saves the token at the client
}
