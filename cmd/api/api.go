package api

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/server/episode"
	"github.com/quanbin27/ReelPlay/server/movie"
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
	subrouter := e.Group("/api/v1")
	userStore := user.NewStore(s.db)
	movieStore := movie.NewStore(s.db)
	episodeStore := episode.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	movieHandler := movie.NewHandler(movieStore)
	episodeHandler := episode.NewHandler(episodeStore, userStore)
	episodeHandler.RegisterRoutes(subrouter)
	userHandler.RegisterRoutes(subrouter)
	movieHandler.RegisterRoutes(subrouter)
	log.Println("Listening on: ", s.address)
	return http.ListenAndServe(s.address, e)
}
