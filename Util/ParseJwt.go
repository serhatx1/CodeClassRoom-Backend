package Util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ParseUserIDFromToken(tokenString string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	UserID, ok := claims["id"].(string)
	if !ok {
		return "", errors.New("user ID not found in claims")
	}

	return UserID, nil
}
