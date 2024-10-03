package middleware

import (
	"eSchool/Service"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		fmt.Println("authHeader", c.Request().Header.Get("Authorization"))

		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("Invalid Authorization header format")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token format")
		}

		tokenStr := parts[1]
		token, claims, err := Service.ValidateToken(tokenStr)
		if err != nil {
			log.Printf("Error validating token: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		if !token.Valid {
			log.Println("Token is not valid")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		log.Printf("Token validated successfully for user: %v", claims["id"])
		c.Set("user", token)
		c.Set("claims", claims)
		return next(c)
	}
}
