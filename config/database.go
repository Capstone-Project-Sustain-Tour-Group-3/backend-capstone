package config

import (
	mysql2 "capstone/drivers/mysql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func LoadDb() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%s",
		ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_NAME, "utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(mysql2.User{})

	DB = db
}
