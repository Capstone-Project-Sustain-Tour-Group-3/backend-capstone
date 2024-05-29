package config

import (
	"fmt"
	"log"

	mysql2 "capstone/drivers/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDb() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%s",
		ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_NAME, "utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
  
	if err = db.AutoMigrate(
		mysql2.User{},
		mysql2.Destination{},
		mysql2.Category{},
		mysql2.Facility{},
		mysql2.DestinationCategory{},
		mysql2.DestinationFacility{},
		mysql2.DestinationAddress{},
		mysql2.DestinationMedia{},
		mysql2.Province{},
    mysql2.Admin{},
	); err != nil {
		log.Fatal(err)
	}

	DB = db
}
