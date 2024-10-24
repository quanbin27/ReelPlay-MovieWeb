package view_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/types"
	"net/http"
)

type Handler struct {
	userStore types.UserStore
}

func NewHandler(userStore types.UserStore) *Handler {
	return &Handler{userStore}
}
func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.Static("/", "templates/")
	e.File("/", "templates/index.html")
	e.File("/index", "templates/index.html")
	e.File("/signin", "templates/signin.html")
	e.File("/signup", "templates/signup.html")
	e.File("/profile", "templates/profile.html")
	e.File("/pricing", "templates/pricing.html")
	e.File("/forgot", "templates/forgot.html")
	e.File("/about", "templates/about.html")
	e.File("/error", "templates/404.html")
	e.File("/faq", "templates/faq.html")
	e.File("/contacts", "templates/contacts.html")
	e.File("/privacy", "templates/privacy.html")
	e.File("/detail", "templates/details.html")
	e.File("/reset-password", "templates/resetpass.html")
	//e.GET("/validate-token", validate, auth.WithJWTAuth(h.userStore))

	e.GET("/watch", func(c echo.Context) error {
		println("Voa day")
		return c.File("templates/watch.html")
	})
	e.File("/search", "templates/search-results.html")
	a := e.Group("/admin")
	a.Static("/", "templates/admin/")
	a.File("/index", "templates/admin/index.html")
	a.File("/movie", "templates/admin/catalog.html")
	a.File("/actor", "templates/admin/actor.html")
	a.File("/add-actor", "templates/admin/add-actor.html")
	a.File("/edit-actor", "templates/admin/edit-actor.html")
	a.File("/director", "templates/admin/director.html")
	a.File("/add-director", "templates/admin/add-director.html")
	a.File("/edit-director", "templates/admin/edit-director.html")
	a.File("/user", "templates/admin/users.html")
	a.File("/edit-user", "templates/admin/edit-user.html")
	a.File("/episode", "templates/admin/episode.html")
	a.File("/add-episode", "templates/admin/add-episode.html")
	a.File("/edit-episode", "templates/admin/edit-episode.html")
}
func validate(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "valid"})
}
