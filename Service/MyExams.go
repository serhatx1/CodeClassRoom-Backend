package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func MyExams(c echo.Context) error {
	var userID string
	var classIDs []uint
	var exams []Models.Exam
	var attendedExams []Models.Exam
	var notAttendedExams []Models.Exam
	var passedExams []Models.Exam

	if err := GetUserID(&userID, c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := DB.DB().
		Table("student_classes").
		Select("class_id").
		Where("users_id = ?", userID).
		Find(&classIDs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if err := DB.DB().
		Where("class_id IN ?", classIDs).
		Find(&exams).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var examIDs []uint
	for _, exam := range exams {
		examIDs = append(examIDs, exam.ID)
	}

	if err := DB.DB().
		Table("exam_students").
		Select("exam_id").
		Where("student_id = ? AND exam_id IN ?", userID, examIDs).
		Find(&attendedExams).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	for _, exam := range exams {
		if time.Now().After(exam.EndOfTheStart) {
			passedExams = append(passedExams, exam)
		} else if contains(attendedExams, exam.ID) {
			attendedExams = append(attendedExams, exam)
		} else {
			notAttendedExams = append(notAttendedExams, exam)
		}

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"attendedExams":    attendedExams,
		"notAttendedExams": notAttendedExams,
		"passedExams":      passedExams,
	})
}

func contains(attendedExams []Models.Exam, examID uint) bool {
	for _, exam := range attendedExams {
		if exam.ID == examID {
			return true
		}
	}
	return false
}
