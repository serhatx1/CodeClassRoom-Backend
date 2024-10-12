package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllProblems(c echo.Context) error {
	var problems []Models.Problem
	var alreadyAdded []Models.Problem
	var notAdded []Models.Problem
	var user Models.Users
	var userID string

	examID := c.QueryParam("examID")
	var questions []Models.Question
	if err := DB.DB().Where("exam_id = ?", examID).Find(&questions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve questions"})
	}

	if err := DB.DB().Preload("TestCases").Find(&problems).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve problems"})
	}

	if err := GetUserID(&userID, c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if len(userID) < TokenLength {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidUserID})
	}

	if err := DB.DB().Where("id = ?", userID).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if user.Role != "teacher" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": InvalidUserRole})
	}

	var exam Models.Exam
	if err := DB.DB().Where("id = ?", examID).Preload("Questions").First(&exam).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	var class Models.Class
	if err := DB.DB().Where("id = ? AND owner_id = ?", exam.ClassID, user.ID).First(&class).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	addedProblems := make(map[uint]struct{})
	for _, question := range questions {
		addedProblems[question.ProblemID] = struct{}{}
	}

	for _, problem := range problems {
		if _, exists := addedProblems[problem.ID]; exists {
			alreadyAdded = append(alreadyAdded, problem)
		} else {
			notAdded = append(notAdded, problem)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"alreadyAdded": alreadyAdded,
		"notAdded":     notAdded,
	})
}
