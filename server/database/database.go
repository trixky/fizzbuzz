// Package gives the necessary tools to initialize everything related to database.
package database

import (
	"fizzbuzz.com/v1/models"
	redis "github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB = nil
var Redis *redis.Client = nil

// Init_postgres initializes postgres.
func Init_postgres(dsn string) (*gorm.DB, error) {
	postgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Postgres = postgres

	if err == nil {
		err = Postgres.AutoMigrate(&models.Api_users{})
	}

	if err == nil {
		err = Postgres.AutoMigrate(&models.Fizzbuzz_request_statistics{})
	}

	return Postgres, err
}

// Init_redis initializes redis.
func Init_redis(addr string, password string, db int) (*redis.Client, error) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := Redis.Ping().Result()

	return Redis, err
}
