package Models

import "gorm.io/gorm"

type ExamStudents struct {
	gorm.Model
	Attended  bool
	Code      string
	ExamID    uint `gorm:"foreignKey:ExamID"`
	StudentID uint `gorm:"foreignKey:studentID"`
}
