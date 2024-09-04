package DB

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/eCode?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err, "Caused in database connection")
		return db, err
	}
	return db, nil
}
