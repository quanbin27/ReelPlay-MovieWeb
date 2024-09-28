package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewMySQLStorage(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
