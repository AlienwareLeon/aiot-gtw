package tools

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)
var db *gorm.DB
func NewConnection() *gorm.DB {
	var dbUri string
	if DatabaseSetting.Type == "mysql" {
		dbUri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			DatabaseSetting.User,
			DatabaseSetting.Password,
			DatabaseSetting.Host,
			DatabaseSetting.Port,
			DatabaseSetting.Name)
	}
	conn, err := gorm.Open(DatabaseSetting.Type, dbUri)
	if err != nil {
		log.Print(err.Error())
	}
	return conn
}

func SetUpConnection() {
	db = NewConnection()
	db.DB().SetMaxIdleConns(10)                   //最大空闲连接数
	db.DB().SetMaxOpenConns(30)                   //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时
	db.SingularTable(true)
	db.LogMode(true)
}

func GetDB() *gorm.DB {
	if err := db.DB().Ping(); err != nil {
		db.Close()
		db = NewConnection()
	}
	return db
}