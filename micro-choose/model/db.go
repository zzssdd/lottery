package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"os"
	"time"
)

var (
	Db *gorm.DB
)

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func LoadMysqlConf() (Conf *MysqlConf) {
	file, err := os.Open("conf/mysql_conf.json")
	if err != nil {
		panic(err)
	}
	byte_data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byte_data, &Conf)
	if err != nil {
		panic(err)
	}
	return
}

func init() {
	conf := LoadMysqlConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	mysql, _ := Db.DB()
	mysql.SetMaxIdleConns(10)
	mysql.SetMaxOpenConns(100)
	mysql.SetConnMaxLifetime(6 * time.Second)
	Db.AutoMigrate(Order{})
}
