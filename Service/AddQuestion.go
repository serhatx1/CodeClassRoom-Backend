package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddQuestion(c echo.Context) error {
	var exam Models.Exam
	var userID string
	var user Models.Users
	var class Models.Class
	var questions []Models.Question
	if err := c.Bind(&exam); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	questions = exam.Questions
	if len(exam.Questions) < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Questions are null"})
	}

	if err := GetUserID(&userID, c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if len(userID) < TokenLength {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidUserID})
	}

	if err := DB.DB().Where("id=?", userID).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if user.Role != "teacher" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": InvalidUserRole})
	}

	if err := DB.DB().Where("id=?", exam.ID).Preload("Questions").First(&exam).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if err := DB.DB().Where("id=? AND owner_id=?", exam.ClassID, user.ID).First(&class).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	for _, question := range questions {
		var existingQuestion Models.Question
		var existingProblem Models.Problem
		if err := DB.DB().Where("problem_id=?", question.ProblemID).First(&existingQuestion).Error; err == nil {
			fmt.Printf("Question with ProblemID %d already exists, skipping...\n", question.ProblemID)
			continue
		}

		if err := DB.DB().Where("id=?", question.ProblemID).First(&existingProblem).Error; err != nil {
			fmt.Printf("Problem with ID %d not found, skipping question...\n", question.ProblemID)
			continue
		}
		question.ProblemID = existingProblem.ID
		question.ExamID = exam.ID
		if err := DB.DB().Save(&question).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "Questions added"})

}
