package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

const (
	ErrInvalidToken = "Invalid user token"
)

func GetUser(c echo.Context) error {
	var User Models.Users
	var userID string

	if err := GetUserID(&userID, c); err != nil {
		return err
	}

	var wg sync.WaitGroup
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = DB.DB().Where("id = ?", userID).First(&User).Error
	}()

	wg.Wait()

	if err != nil {
		return c.String(http.StatusUnprocessableEntity, ErrUserIDNotFound)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"email": User.Email,
		"role":  User.Role,
		"name":  User.Name,
	})
}
