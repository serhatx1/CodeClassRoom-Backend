package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"eSchool/Util"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"unicode"
)

func GenerateUID() string {
	return uuid.NewString()
}
func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case strings.ContainsRune("!@#$%^&*()-_=+[]{}|;:'\",.<>?/", ch):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}

func Register(c echo.Context) error {
	var user *Models.Users
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}
	fmt.Println(user.Email, user.Name)

	err = DB.DB().Where("email = ?", user.Email).First(&user).Error
	if err == nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": "User Already Exists"})
	}

	if !isValidPassword(user.Password) {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": "Password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one digit, and one special character."})
	}

	uid := GenerateUID()
	user.ID = uid
	hashedPassword, err := Util.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}
	user.Password = hashedPassword
	err = DB.DB().Create(user).Error
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}
