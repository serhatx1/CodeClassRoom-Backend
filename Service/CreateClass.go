package Service

import (
	"crypto/rand"
	"eSchool/DB"
	"eSchool/Models"
	"encoding/hex"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GenerateToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

func CreateClass(c echo.Context) error {
	var newClass Models.Class
	var userID string
	if err := c.Bind(&newClass); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid class data"})
	}

	if err := GetUserID(&userID, c); err != nil {
		return err
	}
	if len(userID) < 5 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	newClass.OwnerID = userID
	classToken, err := GenerateToken(16)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate class token"})
	}
	newClass.Token = classToken

	var user Models.Users
	if err := DB.DB().Where("id = ?", userID).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User ID not found"})
	}
	if user.Role == "student" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "User is not teacher	"})
	}
	if err := DB.DB().Create(&newClass).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create class"})
	}

	return c.JSON(http.StatusCreated, newClass)
}
