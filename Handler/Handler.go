package Handler

import (
	middleware "eSchool/Middlewares"
	"eSchool/Service"
	"github.com/labstack/echo/v4"
)

func SetupHandler(e *echo.Echo) {
	e.POST("/api/register", Service.Register)
	e.POST("/api/login", Service.Login)
	e.GET("/api/checkauth", middleware.JWTAuthMiddleware(Service.CheckAuthHandler))

}
