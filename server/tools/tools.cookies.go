package tools

import "net/http"

func Set_session_cookie(session_token string, res http.ResponseWriter) {
	http.SetCookie(res, &http.Cookie{
		Name:  "session",
		Value: session_token,
	}) // saves the token at the client
}

func Remove_session_cookie(res http.ResponseWriter) {
	http.SetCookie(res, &http.Cookie{
		Name:   "session",
		MaxAge: -1,
	})
}
