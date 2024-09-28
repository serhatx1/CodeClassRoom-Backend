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
	e.GET("/api/checkrole", middleware.JWTAuthMiddleware(Service.CheckUserRole))
	e.POST("/api/selectrole", middleware.JWTAuthMiddleware(Service.SelectRole))
	e.POST("/api/class/create", middleware.JWTAuthMiddleware(Service.CreateClass))
	e.GET("/api/class/get", middleware.JWTAuthMiddleware(Service.GetClasses))
	e.GET("/api/class/:id", middleware.JWTAuthMiddleware(Service.CheckParticularClass))

}
