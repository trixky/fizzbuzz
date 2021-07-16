// Package gives the necessary tools to initialize everything related to database.
package database

import (
	"fmt"
	"log"

	"fizzbuzz.com/v1/models"
	"fizzbuzz.com/v1/tools"
	redis "github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB = nil
var Redis *redis.Client = nil

// Init_postgres initializes postgres.
// func Init_postgres(dsn string) *gorm.DB {
func Init_postgres(env tools.Environment) *gorm.DB {
	postgres, err := gorm.Open(postgres.Open("host="+env.Postgres_host+" user="+env.Postgres_user+" password="+env.Postgres_password+" dbname="+env.Postgres_db+" port="+env.Postgres_port+" sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatalln("init postgres: ERROR > ", err)
	}

	Postgres = postgres

	if err := Postgres.AutoMigrate(&models.Api_users{}); err != nil {
		log.Fatalln("init postgres: ERROR > ", err)
	}

	if err = Postgres.AutoMigrate(&models.Fizzbuzz_request_statistics{}); err != nil {
		log.Fatalln("init postgres: ERROR > ", err)
	}

	fmt.Println("init postgres: SUCCESS > ", Postgres)

	return Postgres
}

// Init_redis initializes redis.
func Init_redis() *redis.Client {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	if _, err := Redis.Ping().Result(); err != nil {
		log.Fatalln("init redis: ERROR > ", err)
	}

	fmt.Println("init redis: SUCCESS > ", Redis)

	return Redis
}
