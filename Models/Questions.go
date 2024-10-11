package Models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	ExamID      uint
	ProblemID   uint
	Submissions []Submission
	Problem     Problem `gorm:"foreignKey:ProblemID"`
}
