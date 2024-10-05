package Models

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name     string  `gorm:"not null"`
	Students []Users `gorm:"many2many:student_classes;"`
	Token    string  `gorm:"not null;unique"`
	OwnerID  string  `gorm:"size:36;not null"`
	Owner    *Users  `gorm:"foreignKey:OwnerID;references:ID"`
	Exams    []Exam  `gorm:"foreignKey:ClassID"`
}
