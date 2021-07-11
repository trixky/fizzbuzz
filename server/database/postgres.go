package database

import (
	"fizzbuzz.com/v1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB = nil

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
