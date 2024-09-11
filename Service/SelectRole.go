package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"eSchool/Util"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func SelectRole(c echo.Context) error {
	var user *Models.Users
	var requestBody struct {
		Role string `json:"role"`
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}
	role := requestBody.Role

	if role != "student" && role != "teacher" && role != "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid role. Must be 'student', 'teacher', or empty."})
	}

	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header is missing"})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
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

	if user.Role != "" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "User already has a role and cannot select a new one"})
	}

	user.Role = role
	fmt.Println("role", role)
	if err := DB.DB().Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user role"})
	}
	return c.JSON(http.StatusOK, user)
}
