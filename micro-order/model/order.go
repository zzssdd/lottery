package model

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type Order struct {
	gorm.Model
	Email string `gorm:"type:string;not null"`
	Name  string `gorm:"type:name;not null"`
}

type Prize struct {
	Name        string
	Probability float32
	Num         int
	Pic         string
	CreateTime  string
}

func getRedisName(key string) string {
	return fmt.Sprintf("lottery:prize", key)
}

func OrderList(pageNum int, pageSize int) (orders []*Order, count int64, err error) {
	err = Db.Model(&Order{}).Count(&count).Error
	err = Db.Preload("Pri").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&orders).Error
	return
}

func OrderByEmail(email string, pageNum int, pageSize int) (orders []*Order, count int64, err error) {
	err = Db.Model(&Order{}).Where("Email=?", email).Count(&count).Error
	err = Db.Where("Email=?", email).Preload("Pri").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&orders).Error
	return
}

func OrderAdd(order *Order) error {
	return Db.Create(&order).Error
}

func OrderDel(id int) error {
	return Db.Delete(&Order{}, id).Error
}

func PrizeInfo(name string) (prize *Prize, err error) {
	name = getRedisName(name)
	value, _ := Client.HGetAll(name).Result()
	probability, _ := strconv.ParseFloat(value["probability"], 64)
	num, _ := strconv.Atoi(value["num"])
	prize = &Prize{
		Name:        name,
		Probability: float32(probability),
		Num:         num,
		Pic:         value["pic"],
		CreateTime:  value["createTime"],
	}
	return
}
