// Env contains all the necessary environment variables.
package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/julienschmidt/httprouter"

	"fizzbuzz.com/v1/controllers"
	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/middlewares"
	"fizzbuzz.com/v1/tools"
)

func init() {
	env := tools.Get_environment_variables()
	database.Init_postgres(env)
	database.Init_redis()
}

func main() {
	fmt.Println("\n=================== < START SERVER > ===================")

	// generates the router
	mux := httprouter.New()

	// defines the endpoints
	mux.GET("/fizzbuzz", middlewares.Middleware_token(controllers.Fizzbuzz))
	mux.GET("/stats", middlewares.Middleware_token(controllers.Stats))
	mux.PATCH("/block", middlewares.Middleware_token(controllers.Block))
	mux.POST("/register", controllers.Register)
	mux.POST("/login", controllers.Login)

	// up the server
	log.Fatalln("init http: ERROR > ", http.ListenAndServe(":8080", mux))
}
