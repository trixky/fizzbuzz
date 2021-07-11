package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"net/http"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/routes"
)

var Local_env = Env{}

type Env struct {
	postgres_user     string
	postgres_password string
	postgres_db       string
}

func init_env() error {
	Local_env.postgres_user = os.Getenv("POSTGRES_USER")
	if len(Local_env.postgres_user) < 1 {
		return errors.New("POSTGRES_USER environment variable is not set")
	}
	Local_env.postgres_password = os.Getenv("POSTGRES_PASSWORD")
	if len(Local_env.postgres_password) < 1 {
		return errors.New("POSTGRES_PASSWORD environment variable is not set")
	}
	Local_env.postgres_db = os.Getenv("POSTGRES_DB")
	if len(Local_env.postgres_db) < 1 {
		return errors.New("POSTGRES_DB environment variable is not set")
	}
	return nil
}

func main() {
	fmt.Println("\n\n\n=================== < START SERVER > ===================")

	// ---------------------- INIT ENV
	if err := init_env(); err != nil {
		log.Fatalln("init env: ERROR > ", err)
	} else {
		fmt.Println("init env: SUCCESS")
	}

	// ---------------------- INIT POSTGRES
	if postgres, err := database.Init_postgres("host=postgres user=trixky password=1234 dbname=fizzbuzz port=5432 sslmode=disable"); err != nil {
		log.Fatalln("init postgres: ERROR > ", err)
	} else {
		fmt.Println("init postgres: SUCCESS > ", postgres)
	}

	// ---------------------- INIT REDIS
	if redis, err := database.Init_redis("redis:6379", "", 0); err != nil {
		log.Fatalln("init redis: ERROR > ", err)
	} else {
		fmt.Println("init redis: SUCCESS > ", redis)
	}

	http.HandleFunc("/fizzbuzz", routes.Fizzbuzz)
	http.HandleFunc("/register", routes.Register)
	http.HandleFunc("/login", routes.Login)
	http.HandleFunc("/block", routes.Block)
	http.HandleFunc("/stats", routes.Stats)

	log.Fatalln("init http: ERROR > ", http.ListenAndServe(":8080", nil))
}
