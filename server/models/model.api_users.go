package models

import "gorm.io/gorm"

type Api_users struct {
	gorm.Model
	Pseudo   string `gorm:"pseudo"`
	Password string `gorm:"password"`
	Blocked  bool   `gorm:"blocked"`
}
