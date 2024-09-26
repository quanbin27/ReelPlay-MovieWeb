package api

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/server/user"
	"log"
	"net/http"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
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
