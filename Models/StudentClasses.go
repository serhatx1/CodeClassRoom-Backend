package Models

type StudentClasses struct {
	UsersID string `gorm:"primaryKey;size:36"`
	ClassID uint   `gorm:"primaryKey"`
}
