package models

import "gorm.io/gorm"

type Fizzbuzz_request_statistics struct {
	gorm.Model
	Int1           int
	Int2           int
	Limit          int `gorm:"column:_limit"`
	Str1           string
	Str2           string
	Request_number int `gorm:"default:1"`
}
