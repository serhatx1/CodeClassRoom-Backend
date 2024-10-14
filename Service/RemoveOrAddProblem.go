package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RemoveOrAddProblem(c echo.Context) error {
	var request struct {
		ExamID    uint   `json:"examID"`
		ProblemID uint   `json:"problemID"`
		Action    string `json:"action"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}
	fmt.Println("hi", request.ProblemID, request.ExamID)

	var user Models.Users
	var exam Models.Exam
	var class Models.Class
	var userID string

	if err := GetUserID(&userID, c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized access"})
	}

	if len(userID) < TokenLength {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	if err := DB.DB().Where("id=?", userID).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User not found"})
	}

	if user.Role != "teacher" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Access denied"})
	}

	if err := DB.DB().Where("id=?", request.ExamID).Preload("Questions").First(&exam).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Exam not found"})
	}

	if err := DB.DB().Where("id=? AND owner_id=?", exam.ClassID, user.ID).First(&class).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Class not found"})
	}

	if request.Action != "add" && request.Action != "remove" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid action"})
	}

	if request.Action == "add" {
		newQuestion := Models.Question{
			ProblemID: request.ProblemID,
			ExamID:    request.ExamID,
		}
		if err := DB.DB().Create(&newQuestion).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add problem"})
		}
	} else if request.Action == "remove" {
		if err := DB.DB().Where("problem_id = ? AND exam_id = ?", request.ProblemID, request.ExamID).Delete(&Models.Question{}).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to remove problem"})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "Question processed successfully"})
}
