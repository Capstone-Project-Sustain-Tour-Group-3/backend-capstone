package db

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MockDB() *gorm.DB {
	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatal("error creating mock db : ", err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("error opening mock db : ", err)
	}
	defer db.Close()

	return gormDB
}
