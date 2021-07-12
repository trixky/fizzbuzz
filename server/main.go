package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/julienschmidt/httprouter"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/middlewares"
	"fizzbuzz.com/v1/routes"
)

var Local_env = Env{}

type Env struct {
	postgres_user     string
	postgres_password string
	postgres_db       string
}

func init() {
	fmt.Println("\n\n\n=================== < INIT  SERVER > ===================")

	// ----------------------------------------------- INIT ENV
	Local_env.postgres_user = os.Getenv("POSTGRES_USER")
	if len(Local_env.postgres_user) < 1 {
		log.Fatalln("init env: ERROR > ", "POSTGRES_USER environment variable is not set")
	}
	Local_env.postgres_password = os.Getenv("POSTGRES_PASSWORD")
	if len(Local_env.postgres_password) < 1 {
		log.Fatalln("init env: ERROR > ", "POSTGRES_PASSWORD environment variable is not set")
	}
	Local_env.postgres_db = os.Getenv("POSTGRES_DB")
	if len(Local_env.postgres_db) < 1 {
		log.Fatalln("init env: ERROR > ", "POSTGRES_DB environment variable is not set")
	}

	// ----------------------------------------------- INIT POSTGRES
	if postgres, err := database.Init_postgres("host=postgres user=trixky password=1234 dbname=fizzbuzz port=5432 sslmode=disable"); err != nil {
		log.Fatalln("init postgres: ERROR > ", err)
	} else {
		fmt.Println("init postgres: SUCCESS > ", postgres)
	}

	// ----------------------------------------------- INIT REDIS
	if redis, err := database.Init_redis("redis:6379", "", 0); err != nil {
		log.Fatalln("init redis: ERROR > ", err)
	} else {
		fmt.Println("init redis: SUCCESS > ", redis)
	}
}

func main() {
	fmt.Println("\n=================== < START SERVER > ===================")

	// ----------------------------------------------- START API
	mux := httprouter.New()

	mux.GET("/fizzbuzz", middlewares.Middleware_token(routes.Fizzbuzz))
	mux.GET("/stats", middlewares.Middleware_token(routes.Stats))
	mux.PATCH("/block", middlewares.Middleware_token(routes.Block))
	mux.POST("/register", routes.Register)
	mux.POST("/login", routes.Login)

	log.Fatalln("init http: ERROR > ", http.ListenAndServe(":8080", mux))
}
