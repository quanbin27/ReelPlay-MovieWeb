package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/config"
	"github.com/quanbin27/ReelPlay/types"
	"github.com/quanbin27/ReelPlay/utils"
	"net/http"
	"strconv"
	"time"
)

func CreateJWT(secret []byte, userID int, seconds int64) (string, error) {
	expiration := time.Second * time.Duration(seconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func WithJWTAuth(store types.UserStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Lấy token từ header hoặc query
			tokenString := utils.GetTokenFromRequest(c)
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "unauthorized",
				})
			}

			// Xác thực token
			token, err := ValidateJWT(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
			}
			if !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			}

			// Lấy claims từ token
			claims := token.Claims.(jwt.MapClaims)
			userID, err := strconv.Atoi(claims["user_id"].(string))
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
			}

			// Lấy thông tin người dùng từ store
			u, err := store.GetUserByID(userID)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid user"})
			}

			// Lưu thông tin người dùng vào context
			c.Set("user", &u)

			// Gọi hàm tiếp theo
			return next(c) // Tiếp tục xử lý yêu cầu
		}
	}
}
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}
