package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func SetExam(c echo.Context) error {
	var userID string
	var exam Models.Exam
	var user Models.Users
	var class Models.Class

	if err := c.Bind(&exam); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{ErrInvalidClassData: "Bad request"})
	}

	if exam.ClassID <= 0 {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{ErrInvalidClassData: "Invalid Class ID"})
	}
	if exam.Title == "" {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{ErrInvalidClassData: "Please enter the title of the exam"})
	}
	if exam.DurationOfExam <= 0 {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{ErrInvalidClassData: "Please enter the duration of the exam"})
	}
	if exam.EndOfTheStart.IsZero() {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{ErrInvalidClassData: "Please enter when the exam will start"})
	}

	if err := GetUserID(&userID, c); err != nil {
		return err
	}

	if err := DB.DB().Where("id = ? AND owner_id = ?", exam.ClassID, userID).First(&class).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{ErrInvalidClassData: "Class not found or no access"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	if err := DB.DB().Where("id = ?", userID).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{ErrUserIDNotFound: "User ID not found"})
	}

	if user.Role != "teacher" {
		return c.JSON(http.StatusUnauthorized, map[string]string{InvalidUserRole: "You are not authorized to create an exam"})
	}

	if err := DB.DB().Save(&exam).Error; err != nil {
		return c.JSON(http.StatusBadGateway, echo.ErrBadGateway)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Exam has been created successfully"})
}
