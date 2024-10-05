package Models

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	Attended     bool
	SubmissionID uint
	Passed       bool
}
