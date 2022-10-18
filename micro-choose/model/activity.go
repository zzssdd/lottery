package model

import (
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	gorm.Model
	Cost  int `gorm:"type:int;Default:100"`
	Start time.Time
	End   time.Time
}

func GetActivityInfo(id int) (act *Activity, err error) {
	err = Db.First(&act, id).Error
	return
}
