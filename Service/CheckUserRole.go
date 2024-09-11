package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"eSchool/Util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func CheckUserRole(c echo.Context) error {
	var user Models.Users

	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header is missing"})
	}

	if !strings.HasPrefix(authHeader, "Bearer") {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid authorization header format"})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	userID, err := Util.ParseUserIDFromToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	if err := DB.DB().Where("id = ?", userID).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	switch user.Role {
	case "teacher":
		return c.JSON(http.StatusOK, map[string]string{"message": "User is a teacher", "role": "teacher"})
	case "student":
		return c.JSON(http.StatusOK, map[string]string{"message": "User is a student", "role": "student"})
	default:
		return c.JSON(http.StatusOK, map[string]string{"message": "User has no role", "role": ""})
	}
}
