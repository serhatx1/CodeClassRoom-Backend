package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllProblems(c echo.Context) error {
	var problems []Models.Problem

	if err := DB.DB().Preload("TestCases").Find(&problems).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve problems"})
	}

	return c.JSON(http.StatusOK, problems)
}
