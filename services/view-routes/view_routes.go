package view_routes

import "github.com/labstack/echo/v4"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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
}
