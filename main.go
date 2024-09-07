package main

import (
	"eSchool/DB"
	"eSchool/Handler"
	"eSchool/Models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	e := echo.New()

	// Setup CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Adjust this as needed
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{"*"},
	}))

	// Setup handlers
	Handler.SetupHandler(e)

	// Initialize database
	db := DB.DB()
	if db == nil {
		log.Fatal("Database connection failed")
		return
	}

	// Migrate database schema
	err := db.AutoMigrate(
		&Models.Users{},
		&Models.Class{},
		&Models.Exam{},
		&Models.Question{},
		&Models.Result{},
		&Models.Submission{},
	)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
		return
	}

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
