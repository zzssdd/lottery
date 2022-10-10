package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:string;not null"`
	Password string `gorm:"type:string;not null"`
	Score    int    `gorm:"type:int;Default:0"`
}

func UserRegistry(data *User) error {
	return Db.Create(&data).Error
}

func UserLogin(data *User) bool {
	var count int64
	err := Db.Model(&User{}).Where("Email=? AND Password=?", data.Email, data.Password).Count(&count).Error
	return err == nil && count > 0
}

func UserExist(email string) bool {
	var count int64
	err := Db.Model(&User{}).Where("Email=?", email).Error
	return count > 0 && err == nil
}

func UserList(pageNum int, pageSize int) (count int64, users []*User, err error) {
	err = Db.Model(User{}).Count(&count).Error
	err = Db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users).Error
	return
}

func UserDel(id int) error {
	return Db.Delete(&User{}, id).Error
}
