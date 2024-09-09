package Models

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name     string  `gorm:"notnull"`
	Students []Users `gorm:"many2many:class_students;"`
	Token    string
	OwnerID  string `gorm:"notnull;size:36"`
	Owner    *Users `gorm:"foreignKey:OwnerID;references:ID"`
}
