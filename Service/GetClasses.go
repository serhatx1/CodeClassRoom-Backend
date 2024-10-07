package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetClasses(c echo.Context) error {
	var userID string
	var Classes []Models.Class

	if err := GetUserID(&userID, c); err != nil {
		return err
	}
	if len(userID) < TokenLength {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	err := DB.DB().Where("owner_id=?", userID).Preload("Students").Find(&Classes).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, Classes)
}
