package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"strings"
)

var Validate = validator.New()

func GetTokenFromRequest(c echo.Context) string {
	tokenAuth := c.Request().Header.Get("Authorization")
	if strings.HasPrefix(tokenAuth, "Bearer ") {
		return strings.TrimPrefix(tokenAuth, "Bearer ")
	}

	tokenQuery := c.QueryParam("token")
	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}
