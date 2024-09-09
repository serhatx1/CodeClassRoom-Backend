package Service

import (
	"crypto/rand"
	"eSchool/DB"
	"eSchool/Models"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
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

	if err := c.Bind(&newClass); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid class data"})
	}

	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header is missing"})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid authorization header format"})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	secretKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token claims"})
	}

	UserID, ok := claims["id"].(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User ID not found in claims"})
	}

	newClass.OwnerID = UserID

	classToken, err := GenerateToken(16)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate class token"})
	}
	newClass.Token = classToken

	var user Models.Users
	if err := DB.DB().Where("id = ?", UserID).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User ID not found"})
	}

	if err := DB.DB().Create(&newClass).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create class"})
	}

	return c.JSON(http.StatusCreated, newClass)
}
