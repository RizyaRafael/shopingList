package model

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name string
	Price uint64
	quantity uint64
}

func (Products) TableName() string{
	return "Products"
}