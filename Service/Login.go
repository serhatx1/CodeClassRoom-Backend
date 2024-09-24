package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"eSchool/Util"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func Login(c echo.Context) error {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid req"})
	}

	var user Models.Users
	if err := DB.DB().Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Sorry, we couldn't find an account with that email"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Sorry, this password isn't correct"})
	}

	tokenString, err := Util.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Sorry, Something went wrong"})
	}

	if user.Role == "" {
		return c.JSON(http.StatusAccepted, map[string]string{"token": tokenString})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
