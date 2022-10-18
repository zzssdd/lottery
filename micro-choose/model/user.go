package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:string;not null"`
	Password string `gorm:"type:string;not null"`
	Score    int    `gorm:"type:int;Default:100"`
}

func Check(email string, cost int) bool {
	var user User
	err := Db.Select("Score").Where("Email=?", email).First(&user).Error
	if err != nil {
		return false
	}
	if user.Score < cost {
		return false
	}
	err = Db.Model(&User{}).Where("Email=?", email).UpdateColumn("Score", gorm.Expr("Score-?", cost)).Error
	if err != nil {
		return false
	}
	return true

}

func Recover(email string, cost int) {
	Db.Model(&User{}).Where("Email=?", email).UpdateColumn("Score", gorm.Expr("Score+?", cost))
}
