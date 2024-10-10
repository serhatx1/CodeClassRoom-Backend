package DB

import (
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB
var once sync.Once

func DB() *gorm.DB {
	once.Do(func() {
		dsn := "root:root@tcp(127.0.0.1:3306)/eCode?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		dbInstance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		sqlDB, err := dbInstance.DB()
		if err != nil {
			log.Fatalf("failed to get database instance: %v", err)
		}

		// Connection pool settings
		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(60 * time.Minute)
	})

	return dbInstance
}
