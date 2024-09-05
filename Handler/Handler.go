package Handler

import (
	"eSchool/Service"
	"fmt"
	"github.com/labstack/echo/v4"
)

func SetupHandler(e *echo.Echo) {
	fmt.Println("Setup is working..")
	e.POST("/api/register", Service.Register)

}
