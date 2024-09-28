package api

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/server/user"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type APIServer struct {
	address string
	db      *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		address: addr,
		db:      db,
	}
}
func (s *APIServer) Run() error {
	e := echo.New()
	subrouter := e.Group("/api/vi")
	userStore := user.NewStore(s.db)
	userHander := user.NewHandler(userStore)
	userHander.RegisterRoutes(subrouter)
	log.Println("Listening on: ", s.address)
	return http.ListenAndServe(s.address, e)
}
