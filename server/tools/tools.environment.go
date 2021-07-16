package tools

import (
	"fmt"
	"log"
	"os"
)

// Env contains all the necessary environment variables.
type Environment struct {
	Postgres_user     string
	Postgres_host     string
	Postgres_password string
	Postgres_db       string
	Postgres_port     string
}

// Get_environment_variables reads the environment variables
func Get_environment_variables() Environment {
	fmt.Println("\n\n\n=================== < INIT  SERVER > ===================")

	var env Environment = Environment{}

	env.Postgres_user = os.Getenv("POSTGRES_USER")
	if len(env.Postgres_user) < 1 {
		// If POSTGRES_USER environment is not set
		log.Fatalln("init env: ERROR > ", "POSTGRES_USER environment variable is not set")
	}
	env.Postgres_host = os.Getenv("POSTGRES_HOST")
	if len(env.Postgres_host) < 1 {
		// If POSTGRES_HOST environment is not set
		log.Fatalln("init env: ERROR > ", "POSTGRES_HOST environment variable is not set")
	}
	env.Postgres_password = os.Getenv("POSTGRES_PASSWORD")
	if len(env.Postgres_password) < 1 {
		// If POSTGRES_PASSWORD environment is not set
		log.Fatalln("init env: ERROR > ", "POSTGRES_PASSWORD environment variable is not set")
	}
	env.Postgres_db = os.Getenv("POSTGRES_DB")
	if len(env.Postgres_db) < 1 {
		// If POSTGRES_DB environment is not set
		log.Fatalln("init env: ERROR > ", "POSTGRES_DB environment variable is not set")
	}
	env.Postgres_port = os.Getenv("POSTGRES_PORT")
	if len(env.Postgres_port) < 1 {
		// If POSTGRES_PORT environment is not set
		log.Fatalln("init env: ERROR > ", "POSTGRES_PORT environment variable is not set")
	}

	return env
}
