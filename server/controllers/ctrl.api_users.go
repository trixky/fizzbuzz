package controllers

import (
	"encoding/json"
	"net/http"

	"fizzbuzz.com/v1/extractors"
	"fizzbuzz.com/v1/middlewares"
	"fizzbuzz.com/v1/repositories"
	"fizzbuzz.com/v1/tools"
	"github.com/julienschmidt/httprouter"
)

// Login handles the [POST /login] request.
func Login(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	// #extraction
	extractor := extractors.Login{}
	if err := extractor.Extracts(req); err != nil {
		// If extraction failed.
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(tools.Data_error{Error: err.Error()})
		return
	}

	// #repository [log the api user and get his session token]
	api_user, session_token, err := repositories.Login_api_users(extractor.Pseudo, extractor.Password)
	if err != nil {
		// If repository failed.
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(tools.Data_error{Error: "internal server error"})
		return
	} else if api_user == nil {
		// If pseudo or/and password wrong.
		tools.Remove_session_cookie(res) // remove the potential old token from the client
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(tools.Data_error{Error: "bad pseudo or/and password"})
		return
	} else if api_user.Blocked {
		// If user is blocked.
		tools.Remove_session_cookie(res) // remove the potential old token from the client
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(tools.Data_error{Error: "your account is blocked"})
		return
	}

	// $ success response $
	tools.Set_session_cookie(session_token, res) // saves the token at the client
}

// Register handles the [POST /register] request.
func Register(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	// #extraction
	extractor := extractors.Register{}
	if err := extractor.Extracts(req); err != nil {
		// If extraction failed.
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(tools.Data_error{Error: err.Error()})
		return
	}

	// #repository [register the api user]
	if pseudo_already_taken, err := repositories.Register_api_users(extractor.Pseudo, extractor.Password); err != nil {
		// If repository failed.
		json.NewEncoder(res).Encode(tools.Data_error{Error: "internal server error"})
		return
	} else if pseudo_already_taken {
		// If a user with this pseudo already exists
		res.WriteHeader(http.StatusConflict)
		json.NewEncoder(res).Encode(tools.Data_error{Error: "a user with this pseudo already exists"})
		return
	}

	// $ success response $
}

// Block handles the [PATCH /block] request.
func Block(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	pseudo, _ := req.Context().Value(middlewares.Key_middleware_infos).(*middlewares.Middleware_infos).Get("pseudo") // ignore err

	// #extraction
	extractor := extractors.Block{}
	if err := extractor.Extracts(req); err != nil {
		// If extraction failed.
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(tools.Data_error{Error: err.Error()})
		return
	}

	// #repository [is the user admin ?]
	if is_admin, err := repositories.Is_admin_api_users(pseudo); err != nil {
		// If repository failed.
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(tools.Data_error{Error: "internal server error"})
		return
	} else if !is_admin {
		// If user don't have the required privileges.
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(tools.Data_error{Error: "you don't have the required privileges"})
		return
	}

	// #repository [block the api user]
	if user_not_found, err := repositories.Block_api_users(extractor.Pseudo, extractor.Block_status); err != nil {
		// If repository failed.
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(tools.Data_error{Error: "internal server error"})
		return
	} else if user_not_found {
		// If user don't have the required privileges.
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(tools.Data_error{Error: "no user found with this pseudo"})
		return
	}

	// $ success response $
}
