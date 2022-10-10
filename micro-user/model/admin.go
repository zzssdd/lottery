package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `gorm:"type:string;not null"`
	Password string `gorm:"type:string;not null"`
}

func AdminLogin(data *Admin) bool {
	var count int64
	err := Db.Model(Admin{}).Where("Username=? AND Password=?", data.Username, data.Password).Count(&count).Error
	return err == nil && count > 0
}
