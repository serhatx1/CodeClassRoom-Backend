package Models

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name     string  `gorm:"notnull"`
	Students []Users `gorm:"many2many:class_students;"`
	OwnerID  uint
	Owner    Users `gorm:"foreignKey:OwnerID;references:ID"`
}
