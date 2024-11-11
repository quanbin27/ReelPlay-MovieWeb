package main

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
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
	GgClientID := config.Envs.GgClientID
	GgClientSecret := config.Envs.GgClientSecret
	GgClientCallBackURL := config.Envs.GgClientCallBackURL
	gothic.Store = sessions.NewCookieStore([]byte(config.Envs.JWTSecret))
	goth.UseProviders(
		google.New(GgClientID, GgClientSecret, GgClientCallBackURL, "profile", "email"))
	initStorage(db)
	db.AutoMigrate(&types.User{}, types.Movie{}, types.Country{}, types.Episode{}, types.Bookmark{}, types.Comment{}, types.Rate{}, types.Director{}, types.Actor{}, types.Category{}, types.UserWatched{}, types.CategoryFit{}, types.Role{})
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
