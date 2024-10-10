package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllExams(c echo.Context) error {
	var userID string
	var exams []Models.Exam
	var classes []Models.Class
	var classIDs []uint
	if err := GetUserID(&userID, c); err != nil {
		return err
	}
	if len(userID) < TokenLength {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": InvalidUserID})
	}
	err := DB.DB().Where("owner_id = ?", userID).Find(&classes).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrInvalidClassData})
	}
	if len(classes) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No classes found for this user"})
	}
	for _, class := range classes {
		classIDs = append(classIDs, class.ID)
	}
	err = DB.DB().Where("class_id IN (?)", classIDs).Find(&exams).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch exams"})
	}
	if len(exams) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No exams found for the provided classes"})
	}
	return c.JSON(http.StatusOK, exams)
}
