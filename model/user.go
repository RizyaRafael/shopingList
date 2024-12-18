package model

import (
	"shopingList/handler"

	"gorm.io/gorm"
)

type Users struct {
	ID       uint       `gorm:"primaryKey"`
	Username string     `gorm:"unique"`
	Email    string     `gorm:"unique"`
	Password string     `gorm:"not null"`
	Products []Products `gorm:"foreignKey:UserId"`
}

func (Users) TableName() string {
	return "Users"
}

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPass, err := handler.HashingPass(user.Password)
	user.Password = hashedPass
	return
}
