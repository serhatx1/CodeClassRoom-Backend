package main

import (
	"eSchool/DB"
	"eSchool/Handler"
	"eSchool/Models"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	Handler.SetupHandler(e)
	db := DB.DB()

	err := db.AutoMigrate(&Models.Users{}, &Models.Class{}, &Models.Exam{}, &Models.Question{}, &Models.Result{}, &Models.Submission{})
	if err != nil {
		return
	}

	e.Logger.Fatal(e.Start(":8080"))
}
