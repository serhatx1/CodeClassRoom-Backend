package Models

type Users struct {
	ID       string  `gorm:"primaryKey;size:36"`
	Role     string  `gorm:"not null"`
	Name     string  `gorm:"not null"`
	Email    string  `gorm:"not null;unique"`
	Password string  `gorm:"not null"`
	Classes  []Class `gorm:"many2many:student_classes;"`
}
