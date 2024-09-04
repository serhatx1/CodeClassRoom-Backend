package Models

import "gorm.io/gorm"

type Exam struct {
	gorm.Model
	Title       string
	Description string
	ClassID     uint
	Questions   []Question
}
