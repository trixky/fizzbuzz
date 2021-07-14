// Env contains all the necessary environment variables.
package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/julienschmidt/httprouter"

	"fizzbuzz.com/v1/controllers"
	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/middlewares"
)

// Env contains all the necessary environment variables.
type Env struct {
	postgres_user     string
	postgres_host     string
	postgres_password string
	postgres_db       string
	postgres_port     string
}

func init() {
	fmt.Println("\n\n\n=================== < INIT  SERVER > ===================")

	env := Env{}

	// initialize/read environment variables
	env.postgres_user = os.Getenv("POSTGRES_USER")
	if len(env.postgres_user) < 1 {
		log.Fatalln("init env: ERROR > ", "POSTGRES_USER environment variable is not set")
	}
	env.postgres_host = os.Getenv("POSTGRES_HOST")
	if len(env.postgres_host) < 1 {
		log.Fatalln("init env: ERROR > ", "POSTGRES_HOST environment variable is not set")
	}
	env.postgres_password = os.Getenv("POSTGRES_PASSWORD")
	if len(env.postgres_password) < 1 {
		log.Fatalln("init env: ERROR > ", "POSTGRES_PASSWORD environment variable is not set")
	}
	env.postgres_db = os.Getenv("POSTGRES_DB")
	if len(env.postgres_db) < 1 {
		log.Fatalln("init env: ERROR > ", "POSTGRES_DB environment variable is not set")
	}
	env.postgres_port = os.Getenv("POSTGRES_PORT")
	if len(env.postgres_port) < 1 {
		log.Fatalln("init env: ERROR > ", "POSTGRES_PORT environment variable is not set")
	}

	// initialize postgres
	if postgres, err := database.Init_postgres("host=" + env.postgres_host + " user=" + env.postgres_user + " password=" + env.postgres_password + " dbname=" + env.postgres_db + " port=" + env.postgres_port + " sslmode=disable"); err != nil {
		log.Fatalln("init postgres: ERROR > ", err)
	} else {
		fmt.Println("init postgres: SUCCESS > ", postgres)
	}

	// initialize redis
	if redis, err := database.Init_redis("redis:6379", "", 0); err != nil {
		log.Fatalln("init redis: ERROR > ", err)
	} else {
		fmt.Println("init redis: SUCCESS > ", redis)
	}
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
