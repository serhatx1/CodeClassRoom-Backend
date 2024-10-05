package Models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	Attended   bool
	Code       string
	QuestionID uint
	StudentID  uint
	Results    []Result `gorm:"foreignKey:SubmissionID"`
}
