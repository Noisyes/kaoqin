package common

import (
	"kaoqin/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/kaoqin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database,err :" + err.Error())
	}
	DB = db
	db.AutoMigrate(&user.User{})
	return db
}

func GetDB() *gorm.DB {
	return DB
}
