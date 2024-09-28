package main

import (
	"github.com/quanbin27/ReelPlay/cmd/api"
	"github.com/quanbin27/ReelPlay/config"
	"github.com/quanbin27/ReelPlay/db"
	"github.com/quanbin27/ReelPlay/types"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := config.Envs.DSN
	db, err := db.NewMySQLStorage(dsn)
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)
	db.AutoMigrate(&types.User{})
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
func initStorage(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")
}
