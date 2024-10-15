package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/config"
	"github.com/quanbin27/ReelPlay/services/auth"
	"github.com/quanbin27/ReelPlay/types"
	"github.com/quanbin27/ReelPlay/utils"
	"net/http"
	"strconv"
)

type Handler struct {
	store        types.UserStore
	emailService types.EmailService
}

func NewHandler(store types.UserStore, emailService types.EmailService) *Handler {
	return &Handler{store, emailService}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {
	e.POST("/login", h.handleLogin)
	e.POST("/register", h.handleRegister)
	e.POST("/forgot-password", h.ForgotPassword)
	e.POST("/reset-password", h.ResetPassword)
}
func (h *Handler) ResetPassword(c echo.Context) error {
	req := struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	token, err := auth.ValidateJWT(req.Token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	if !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
	}
	claims := token.Claims.(jwt.MapClaims)
	// Hash mật khẩu mới
	hashedPassword, err := auth.HashPassword(req.NewPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}
	userID, err := strconv.Atoi(claims["user_id"].(string))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	err = h.store.UpdatePassword(userID, hashedPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update password"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Password updated successfully"})
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
	token, err := auth.CreateJWT(secret, u.ID, config.Envs.JWTExpirationInSeconds)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal services error"})
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal services error"})
	}
	user := types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	}
	err = h.store.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal services error"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
func (h *Handler) ForgotPassword(c echo.Context) error {
	req := struct {
		Email string `json:"email"`
	}{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	user, err := h.store.GetUserByEmail(req.Email)
	if err != nil || user == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email not found"})
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, user.ID, 3600)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	resetLink := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", token)
	print(resetLink)
	// Gửi email với link reset mật khẩu
	err = h.emailService.SendResetPasswordEmail(req.Email, resetLink)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send email"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Password reset email sent"})
}
