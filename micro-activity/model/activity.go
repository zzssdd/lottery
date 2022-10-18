package model

import (
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	gorm.Model
	Name  string `gorm:"type:string;not null"`
	Desc  string `gorm:"type:string;not null"`
	Pic   string `gorm:"type:string;not null"`
	User  string `gorm:"type:string;not null"`
	Cost  int    `gorm:"type:int;Default:100"`
	Type  int    `gorm:"type:int;Default:0"`
	Start time.Time
	End   time.Time
}

func ActiveAdd(active *Activity) (error, int) {
	return Db.Create(&active).Error, int(active.ID)
}

func ActiveDel(id int) error {
	err := Db.Delete(&Activity{}, id).Error
	if err != nil {
		return err
	}
	return ActiveDelPrizes(id)
}

func ActiveEdit(active *Activity, id int) error {
	return Db.Model(&Activity{}).Where("ID=?", id).Updates(&active).Error
}

func ActiveList(pageNum, pageSize int) (actives []*Activity, count int64, err error) {
	err = Db.Model(&Activity{}).Count(&count).Error
	err = Db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&actives).Error
	return
}

func ActiveInfo(id int) (active *Activity, err error) {
	err = Db.Where("ID=?", id).First(&active).Error
	return
}
