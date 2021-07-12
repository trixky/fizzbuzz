package models

import "gorm.io/gorm"

type Fizzbuzz_request_statistics struct {
	gorm.Model
	Int1           int    `gorm:"not null"`
	Int2           int    `gorm:"not null"`
	Limit          int    `gorm:"not null;column:_limit"`
	Str1           string `gorm:"not null"`
	Str2           string `gorm:"not null"`
	Request_number int    `gorm:"not null;default:1"`
}
