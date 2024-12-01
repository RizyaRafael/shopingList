package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
}

func (Users) TableName() string {
	return "Users"
}
