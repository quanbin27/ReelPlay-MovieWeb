package user

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/types"
	"net/http"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {
	e.POST("/login", h.handleLogin)
	e.POST("/register", h.handleRegister)
}
func (h *Handler) handleLogin(c echo.Context) error {
	return nil
}
func (h *Handler) handleRegister(c echo.Context) error {
	var payload types.RegisterUserPayLoad
	if err := c.Bind(payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		return c.String(http.StatusBadRequest, "email already taken")
	}
	return nil
}
