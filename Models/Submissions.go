package Models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	Code       string
	QuestionID uint
	StudentID  uint
	Results    []Result
}
