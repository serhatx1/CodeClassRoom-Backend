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

	r := DB.DB().Where("email = ?", user.Email).First(&user)

	if r.RowsAffected > 0 {
		return c.JSON(http.StatusUnprocessableEntity, fmt.Errorf("User Already Exists").Error())
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
