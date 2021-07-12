package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/json_struct"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

type Key string

const Key_middleware_infos Key = "middleware_infos"

type Middleware_infos struct {
	token   string
	pseudo  string
	admin   string
	blocked string
}

func (mi Middleware_infos) Get_pseudo(info string) (string, error) {
	switch info {
	case "token":
		return mi.token, nil
	case "pseudo":
		return mi.pseudo, nil
	case "admin":
		return mi.admin, nil
	case "blocked":
		return mi.blocked, nil
	default:
		return "", errors.New("[" + info + "] key is invalid")
	}
}

func Middleware_token(toto func(res http.ResponseWriter, req *http.Request, ps httprouter.Params)) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		// get token from cookie
		cookie, err := req.Cookie("session")

		// check we have the token
		if err != nil {
			res.WriteHeader(http.StatusUnauthorized)
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "'session' cookie is corrupted or missing"})
			return
		}

		token := cookie.Value

		// check token format (is uuid ?)
		if _, err = uuid.FromString(token); err != nil {
			res.WriteHeader(http.StatusUnauthorized)
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "'session' cookie is corrupted"})
			return

		}

		// check token validity from redis
		result := database.Redis.Get("token>" + token)
		if result.Err() != nil {
			res.WriteHeader(http.StatusUnauthorized)
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "your token from 'session' cookie is not valid"})
			return
		}

		//
		ctx := context.WithValue(req.Context(), Key_middleware_infos, &Middleware_infos{pseudo: result.Val(), token: token})

		toto(res, req.WithContext(ctx), ps)
	}
}
