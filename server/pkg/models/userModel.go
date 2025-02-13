package models

import "gorm.io/gorm"

type UserTrainers struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Username string
	Password string
}
