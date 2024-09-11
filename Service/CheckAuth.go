package Service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
)

func ValidateToken(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("failed to validate token: %w", err)
	}

	return token, claims, nil
}
