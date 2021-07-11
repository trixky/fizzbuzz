package routes

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/json_struct"
	"fizzbuzz.com/v1/models"
	uuid "github.com/satori/go.uuid"

	"gorm.io/gorm"
)

func Login(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/login" || req.Method != "POST" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	pseudo := req.FormValue("pseudo")
	password := req.FormValue("password")

	// check if we have pseudo and password
	if len(pseudo) < 1 || len(password) < 1 {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "pseudo or password missing/bad formatted"})
		return
	}

	// hash the password
	password_hasheur := md5.New()
	password_hasheur.Write([]byte(password))
	hashed_password := password_hasheur.Sum(nil)
	api_user := models.Api_users{}

	if err := database.Postgres.Table("api_users").Where(map[string]interface{}{"pseudo": pseudo, "password": hex.EncodeToString(hashed_password)}).First(&api_user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "bad pseudo or/and password"})
		} else {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		}
		return
	}

	if api_user.Blocked {
		http.SetCookie(res, &http.Cookie{
			Name:   "session",
			MaxAge: -1,
		})
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "your account is blocked"})
		return
	}

	new_session_token := uuid.NewV4().String()
	if database.Redis.Append("tokens>"+pseudo, "|"+new_session_token).Err() != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	}
	if database.Redis.Set("token>"+new_session_token, pseudo, 0).Err() != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
		return
	}
	http.SetCookie(res, &http.Cookie{
		Name:  "session",
		Value: new_session_token,
	})
}
