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
	e.GET("/user", h.SearchUsersHandler, auth.WithJWTAdminAuth(h.store))
	e.DELETE("/user/:id", h.DeleteUserHandler, auth.WithJWTAdminAuth(h.store))
	e.PUT("/user/:id/info", h.UpdateUserInfoHandler, auth.WithJWTAdminAuth(h.store))
	e.PUT("/user/:id/password", h.UpdateUserPasswordHandler, auth.WithJWTAdminAuth(h.store))
	e.GET("/user/:id", h.GetUserByIDHandler, auth.WithJWTAdminAuth(h.store))
	e.PUT("/user/:id/unlock", h.UnlockUserHandler, auth.WithJWTAdminAuth(h.store))
}

// Trong file handler.go hoặc tương tự
func (h *Handler) UnlockUserHandler(c echo.Context) error {
	// Lấy userID từ param
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Gọi hàm khôi phục người dùng từ Store
	if err := h.store.UnlockUser(userID); err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error unlocking user"})
	}

	// Trả về phản hồi thành công
	return c.JSON(http.StatusOK, map[string]string{"message": "User unlocked successfully"})
}

func (h *Handler) GetUserByIDHandler(c echo.Context) error {
	// Lấy userID từ param
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Gọi hàm lấy thông tin người dùng từ Store
	user, err := h.store.GetUserByID(userID)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error retrieving user"})
	}

	// Trả về thông tin người dùng
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUserInfoHandler(c echo.Context) error {
	// Lấy userID từ param
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Parse dữ liệu JSON từ request body
	updatedData := make(map[string]interface{})
	if err := c.Bind(&updatedData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	// Gọi hàm cập nhật thông tin người dùng từ Store
	err = h.store.UpdateUserInfo(userID, updatedData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error updating user"})
	}

	// Trả về phản hồi thành công
	return c.JSON(http.StatusOK, map[string]string{"message": "User info updated successfully"})
}

func (h *Handler) UpdateUserPasswordHandler(c echo.Context) error {
	// Lấy userID từ param
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Parse dữ liệu JSON từ request body
	requestData := struct {
		Password string `json:"password"`
	}{}

	if err := c.Bind(&requestData); err != nil || requestData.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data or empty password"})
	}

	// Gọi hàm cập nhật mật khẩu từ Store
	err = h.store.UpdateUserPassword(userID, requestData.Password)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error updating password"})
	}

	// Trả về phản hồi thành công
	return c.JSON(http.StatusOK, map[string]string{"message": "Password updated successfully"})
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
		RoleID:    1,
	}
	err = h.store.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal services error"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
func (h *Handler) DeleteUserHandler(c echo.Context) error {
	// Lấy userID từ param
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Gọi hàm xóa mềm người dùng từ Store
	err = h.store.DeleteUserSoft(userID)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error deleting user"})
	}

	// Trả về phản hồi thành công
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func (h *Handler) SearchUsersHandler(c echo.Context) error {
	// Lấy từ khóa tìm kiếm từ query parameters
	keyword := c.QueryParam("keyword")

	// Lấy các tham số pagination
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 20 // Số bản ghi mặc định mỗi trang là 10
	}

	// Gọi hàm SearchUsers từ store (giả sử bạn đã có đối tượng store trong Handler)
	users, totalRecords, err := h.store.SearchUsers(keyword, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error searching users"})
	}

	// Trả về kết quả dưới dạng JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"total": totalRecords,
		"page":  page,
		"limit": limit,
		"data":  users,
	})
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
	// Gửi email với link reset mật khẩu
	err = h.emailService.SendResetPasswordEmail(req.Email, resetLink)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send email"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Password reset email sent"})
}
