package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckParticularClass(c echo.Context) error {
	classID := c.Param("id")
	var userID string

	if err := GetUserID(&userID, c); err != nil {
		return err
	}
	if len(userID) < 5 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	var class Models.Class
	err := DB.DB().Where("id = ? AND owner_id = ?", classID, userID).First(&class).Error
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Class not found or you do not have access"})
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, class)
}
