// Package models gives access to all the model structures used to exchange with the database.
package models

import "gorm.io/gorm"

// Api_users is the model corresponding to the "api_users" table.
type Api_users struct {
	gorm.Model
	Pseudo   string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Blocked  bool   `gorm:"not null;default:false"`
	Admin    bool   `gorm:"not null;default:false"`
}
