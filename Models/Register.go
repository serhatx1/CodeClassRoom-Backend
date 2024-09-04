package Models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID       int    `gorm:"unique;notnull"`
	Role     string `gorm:"notnull"`
	Name     string `gorm:"notnull"`
	Email    string `gorm:"notnull"`
	Password string `gorm:"notnull"`
}
