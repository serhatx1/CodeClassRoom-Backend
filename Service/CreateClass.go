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

const (
	TokenLength              = 36
	ErrInvalidClassData      = "Invalid class data"
	ErrInvalidUserID         = "Invalid user ID"
	ErrUserIDNotFound        = "User ID not found"
	ErrUserNotTeacher        = "User does not have permission to create a class"
	ErrFailedTokenGeneration = "Failed to generate class token"
	ErrFailedClassCreation   = "Failed to create class"
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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidClassData})
	}

	if err := GetUserID(&userID, c); err != nil {
		return err
	}

	if len(userID) < TokenLength {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidUserID})
	}
	newClass.OwnerID = userID

	classToken, err := GenerateToken(TokenLength)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrFailedTokenGeneration})
	}
	newClass.Token = classToken

	var user Models.Users
	if err := DB.DB().Where("id = ?", userID).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrUserIDNotFound})
	}

	if user.Role != "teacher" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": ErrUserNotTeacher})
	}

	if err := DB.DB().Create(&newClass).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrFailedClassCreation})
	}

	return c.JSON(http.StatusCreated, newClass)
}
