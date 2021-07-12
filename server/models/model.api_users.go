package models

import "gorm.io/gorm"

type Api_users struct {
	gorm.Model
	Pseudo   string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Blocked  bool   `gorm:"not null;default:false"`
	Admin    bool   `gorm:"not null;default:false"`
}
