package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Email string `gorm:"type:string;not null"`
	Name  string `gorm:"type:name;not null"`
}

func OrderAdd(order *Order) error {
	return Db.Create(&order).Error
}
