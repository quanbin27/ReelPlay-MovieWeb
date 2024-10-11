package api

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/config"
	"github.com/quanbin27/ReelPlay/services/actor"
	"github.com/quanbin27/ReelPlay/services/bookmark"
	"github.com/quanbin27/ReelPlay/services/category"
	"github.com/quanbin27/ReelPlay/services/comment"
	"github.com/quanbin27/ReelPlay/services/director"
	"github.com/quanbin27/ReelPlay/services/email"
	"github.com/quanbin27/ReelPlay/services/episode"
	"github.com/quanbin27/ReelPlay/services/movie"
	"github.com/quanbin27/ReelPlay/services/user"
	view_routes "github.com/quanbin27/ReelPlay/services/view-routes"
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
	//e.Static("/static", "templates")
	viewHandler := view_routes.NewHandler()
	viewHandler.RegisterRoutes(e)
	subrouter := e.Group("/api/v1")
	userStore := user.NewStore(s.db)
	movieStore := movie.NewStore(s.db)
	episodeStore := episode.NewStore(s.db)
	categoryStore := category.NewStore(s.db)
	actorStore := actor.NewStore(s.db)
	cmtStore := comment.NewStore(s.db)
	bookmarkStore := bookmark.NewStore(s.db)
	directorStore := director.NewStore(s.db)
	emailService := email.NewEmailService("smtp.gmail.com", 587, config.Envs.EmailUsername, config.Envs.Emailpassword, config.Envs.Emailfrom)
	userHandler := user.NewHandler(userStore, emailService)
	movieHandler := movie.NewHandler(movieStore, categoryStore, actorStore, directorStore)
	episodeHandler := episode.NewHandler(episodeStore, userStore)
	cmtHandler := comment.NewHandler(cmtStore, userStore)
	bookmarkHandler := bookmark.NewHandler(bookmarkStore, userStore)
	bookmarkHandler.RegisterRoutes(subrouter)
	cmtHandler.RegisterRoutes(subrouter)
	episodeHandler.RegisterRoutes(subrouter)
	userHandler.RegisterRoutes(subrouter)
	movieHandler.RegisterRoutes(subrouter)
	log.Println("Listening on: ", s.address)
	return http.ListenAndServe(s.address, e)
}
