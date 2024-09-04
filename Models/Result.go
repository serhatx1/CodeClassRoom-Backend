package Models

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	SubmissionID uint
	Passed       bool
}
