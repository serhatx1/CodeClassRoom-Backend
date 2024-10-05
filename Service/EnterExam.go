package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type ExamObject struct {
	ExamID int `json:"exam_id"`
}

func setExam(c echo.Context) error {
	var userID string
	var user Models.Users
	var examInstance ExamObject
	var examStudents Models.ExamStudents
	var examInfo Models.Exam

	if err := c.Bind(&examInstance); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": "Invalid request data"})
	}

	if err := GetUserID(&userID, c); err != nil {
		return err
	}

	if len(userID) < TokenLength {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid User ID"})
	}

	if err := DB.DB().Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if user.Role != "student" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "User role is not allowed"})
	}

	if err := DB.DB().Where("id = ? AND ClassID IN ?", examInstance.ExamID, CheckUserRole).First(&examInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Exam not found"})
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := DB.DB().Where("ExamID = ? AND StudentID = ?", examInstance.ExamID, userID).First(&examStudents).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if examInfo.EndOfTheStart.Before(time.Now()) {

				return c.JSON(http.StatusOK, map[string]interface{}{"message": "You can attend the exam"})
			}
			return c.JSON(http.StatusNotFound, map[string]string{"error": "You are not registered for this exam"})
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if examStudents.Attended {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "You already had the exam"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Ready to attend the exam"})
}
