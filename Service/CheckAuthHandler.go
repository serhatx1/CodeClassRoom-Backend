package Service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckAuthHandler(c echo.Context) error {
	user := c.Get("user")
	claims := c.Get("claims")
	if user == nil || claims == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Token is valid",
		"user":    claims,
	})
}
