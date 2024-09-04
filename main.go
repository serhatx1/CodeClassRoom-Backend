package main

import (
	"eSchool/DB"
	"eSchool/Handler"
	"eSchool/Models"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	Handler.SetupHandler(e)
	db, err := DB.DB()
	if err != nil {
		fmt.Println(db)
		panic(err)
	}

	err = db.AutoMigrate(&Models.Users{}, &Models.Class{}, &Models.Exam{}, &Models.Question{}, &Models.Result{}, &Models.Login{}, &Models.Submission{})
	if err != nil {
		return
	}

	e.Logger.Fatal(e.Start(":8080"))
}
