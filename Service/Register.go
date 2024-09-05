package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"eSchool/Util"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GenerateUID() string {
	return uuid.NewString()
}
func Register(c echo.Context) error {
	var user *Models.Users
	err := c.Bind(&user)
	fmt.Println(user.Email)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	fmt.Println("user.ID")

	err = DB.DB().First(&Models.Users{}, "email = ?", user.Email).Error
	if err.Error == nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	uid := GenerateUID()
	user.ID = uid
	hashedPassword, err := Util.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	user.Password = hashedPassword
	if len(user.Role) == 0 {
		user.Role = "student"
	}
	err = DB.DB().Create(user).Error
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusCreated, user)
}
