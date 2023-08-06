package configs

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	b, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	db = b
}

func GetDB() *gorm.DB {
	return db
}
