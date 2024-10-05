package Models

import (
	"gorm.io/gorm"
	"time"
)

type Exam struct {
	gorm.Model
	Title          string
	Description    string
	ClassID        uint
	EndOfTheStart  time.Time
	DurationOfExam uint
	Questions      []Question
	Class          Class          `gorm:"foreignKey:ClassID"`
	ExamStudents   []ExamStudents `gorm:"foreignKey:StudentID"`
}
