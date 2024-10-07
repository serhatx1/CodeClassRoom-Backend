package main

import (
	"eSchool/DB"
	"eSchool/Handler"
	"eSchool/Models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Adjust this as needed
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{"*"},
	}))

	Handler.SetupHandler(e)

	// Initialize database
	db := DB.DB()
	if db == nil {
		log.Fatal("Database connection failed")
		return
	}

	err := db.AutoMigrate(
		&Models.Users{},
		&Models.Class{},
		&Models.Result{},
		&Models.Submission{},
		&Models.Question{},
		&Models.Exam{},
		&Models.ExamStudents{},
		&Models.StudentClasses{},
		&Models.Problem{},
		&Models.TestCase{},
	)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
		return
	}

	e.Logger.Fatal(e.Start(":8080"))
}
