package database

import (
	"go-rest/svc"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(dsn string) *gorm.DB {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&svc.Student{})

	return db
}
