// Package middlewares allows to authenticate a user with redis using a token retrieved from the 'session' cookie.
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

// Middleware_infos contains all the information potentially readable in the middleware.
type Middleware_infos struct {
	token   string
	pseudo  string
	admin   string
	blocked string
}

// Get Gives access to informations retrieved by the middleware from the struct Middleware_infos.
func (mi Middleware_infos) Get(info string) (string, error) {
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

// Middleware_token authenticates the user with redis and returns informations he knows in the context.
func Middleware_token(toto func(res http.ResponseWriter, req *http.Request, ps httprouter.Params)) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		cookie, err := req.Cookie("session") // get token from cookie

		if err != nil {
			// If the cookie is corrupted or missing.
			res.WriteHeader(http.StatusUnauthorized)
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "'session' cookie is corrupted or missing"})
			return
		}

		token := cookie.Value

		if _, err = uuid.FromString(token); err != nil {
			// If the cookie format is corrupted.
			res.WriteHeader(http.StatusUnauthorized)
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "'session' cookie format is corrupted"})
			return

		}

		result := database.Redis.Get("token>" + token)
		if result.Err() != nil {
			// If the token from the cookie is not valid.
			res.WriteHeader(http.StatusUnauthorized)
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(json_struct.Data_error{Error: "your token from 'session' cookie is not valid"})
			return
		}

		ctx := context.WithValue(req.Context(), Key_middleware_infos, &Middleware_infos{pseudo: result.Val(), token: token}) // save known informations in the context for the rest

		toto(res, req.WithContext(ctx), ps) // next
	}
}
