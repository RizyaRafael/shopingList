package model

type Products struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Price    uint64
	Quantity uint64
	UserId   uint
	ImageUrl string
}

func (Products) TableName() string {
	return "Products"
}
