package Service

import (
	"eSchool/DB"
	"eSchool/Models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func AddQuestion(c echo.Context) error {
	var exam Models.Exam
	var userID string
	var user Models.Users
	var class Models.Class

	if err := c.Bind(&exam); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	fmt.Println(exam)

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

	var wg sync.WaitGroup
	results := make(chan error, len(exam.Questions))

	for _, question := range exam.Questions {
		wg.Add(1)
		go func(q Models.Question) {
			defer wg.Done()
			var existingQuestion Models.Question
			var existingProblem Models.Problem

			if err := DB.DB().Where("problem_id=?", q.ProblemID).First(&existingQuestion).Error; err == nil {
				fmt.Printf("Question with ProblemID %d already exists, skipping...\n", q.ProblemID)
				results <- nil
				return
			}

			if err := DB.DB().Where("id=?", q.ProblemID).First(&existingProblem).Error; err != nil {
				fmt.Printf("Problem with ID %d not found, skipping question...\n", q.ProblemID)
				results <- nil
				return
			}

			q.ProblemID = existingProblem.ID
			q.ExamID = exam.ID
			if err := DB.DB().Save(&q).Error; err != nil {
				results <- err
				return
			}
			results <- nil
		}(question)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var errorOccurred bool
	for err := range results {
		if err != nil {
			errorOccurred = true
			break
		}
	}

	if errorOccurred {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add one or more questions"})
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "Questions added"})
}
