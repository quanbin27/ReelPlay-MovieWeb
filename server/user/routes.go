package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/config"
	"github.com/quanbin27/ReelPlay/server/auth"
	"github.com/quanbin27/ReelPlay/types"
	"github.com/quanbin27/ReelPlay/utils"
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
	var payload types.LoginUserPayLoad
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": errors[0].Error()})
	}
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "not found, invalid email"})
	}
	if !auth.CheckPassword(u.Password, []byte(payload.Password)) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid password"})
	}
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(c echo.Context) error {
	var payload types.RegisterUserPayLoad
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": errors[0].Error()})
	}
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "email already exists"})
	}
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	user := types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	}
	err = h.store.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
