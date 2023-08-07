package configs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	dsn := "host=localhost user=postgres password=1079 dbname=todoapi port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect database: %v", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
