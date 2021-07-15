// Package models gives access to all the model structures used to exchange with the database.
package models

import "gorm.io/gorm"

// Fizzbuzz_request_statistics is the model corresponding to the "fizzbuzz_request_statistics" table.
type Fizzbuzz_request_statistics struct {
	gorm.Model
	Int1           int    `gorm:"not null"`
	Int2           int    `gorm:"not null"`
	Limit          int    `gorm:"not null;column:_limit"`
	Str1           string `gorm:"not null"`
	Str2           string `gorm:"not null"`
	Request_number int    `gorm:"not null;default:1"`
}
