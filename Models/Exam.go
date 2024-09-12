package Models

import "gorm.io/gorm"

type Exam struct {
	gorm.Model
	Title       string
	Description string
	ClassID     uint // Foreign key to the Class
	Questions   []Question
	Class       Class `gorm:"foreignKey:ClassID"` // Define the relationship
}
