package Models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title       string
	Prompt      string
	ExamID      uint
	Submissions []Submission
}
