package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TokenClass struct {
	Token string `json:"token"`
}

const (
	InvalidUserRole = "Invalid user role, user must be a student"
	InvalidUserID   = "Invalid user ID"
	TokenIsInvalid  = "Token is invalid"
)

func JoinClass(c echo.Context) error {
	fmt.Println("hellos")
	var userID string
	var user Models.Users
	var tokenClass TokenClass
	var class Models.Class

	if err := c.Bind(&tokenClass); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{TokenIsInvalid: err.Error()})
	}
	fmt.Println(tokenClass.Token)
	if err := GetUserID(&userID, c); err != nil {
		return err
	}

	if len(userID) < TokenLength {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": InvalidUserID})
	}

	if err := DB.DB().Where("id=?", userID).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if user.Role != "student" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": InvalidUserRole})
	}
	if err := DB.DB().Where("token=?", tokenClass.Token).Preload("Students").First(&class).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{TokenIsInvalid: err.Error()})
	}
	fmt.Println(class)
	for _, student := range class.Students {
		if student.ID == user.ID {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "User already enrolled in this class"})
		}
	}

	class.Students = append(class.Students, user)

	if err := DB.DB().Save(&class).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"class": class})
}
