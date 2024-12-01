package model

import (
	"log"
	"shopingList/handler"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
}

func (Users) TableName() string {
	return "Users"
}

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPass, err := handler.HashingPass(user.Password)
	log.Print(hashedPass)
	user.Password = hashedPass
	return 
}
