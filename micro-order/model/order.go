package model

import (
	"gorm.io/gorm"
	"strconv"
)

type Order struct {
	gorm.Model
	Email string `gorm:"type:string;not null"`
	Name  string `gorm:"type:string;not null"`
}

type Prize struct {
	Name       string
	Num        int
	Pic        string
	CreateTime string
}

func OrderList(pageNum int, pageSize int) (orders []*Order, count int64, err error) {
	err = Db.Model(&Order{}).Count(&count).Error
	err = Db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("created_at Desc").Find(&orders).Error
	return
}

func OrderByEmail(email string, pageNum int, pageSize int) (orders []*Order, count int64, err error) {
	err = Db.Model(&Order{}).Where("Email=?", email).Count(&count).Error
	err = Db.Where("Email=?", email).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&orders).Error
	return
}

func OrderAdd(order *Order) error {
	return Db.Create(&order).Error
}

func OrderDel(id int) error {
	return Db.Delete(&Order{}, id).Error
}

func PrizeInfo(name string) (prize *Prize, err error) {
	value, _ := Client.HGetAll(name).Result()
	num, _ := strconv.Atoi(value["num"])
	prize = &Prize{
		Name:       name,
		Num:        num,
		Pic:        value["pic"],
		CreateTime: value["createTime"],
	}
	return
}
