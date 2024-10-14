package Handler

import (
	middleware "eSchool/Middlewares"
	"eSchool/Service"
	"github.com/labstack/echo/v4"
)

func SetupHandler(e *echo.Echo) {
	e.POST("/api/register", Service.Register)
	e.POST("/api/login", Service.Login)
	e.GET("/api/user/get", middleware.JWTAuthMiddleware(Service.GetUser))
	e.GET("/api/checkauth", middleware.JWTAuthMiddleware(Service.CheckAuthHandler))
	e.GET("/api/checkrole", middleware.JWTAuthMiddleware(Service.CheckUserRole))
	e.POST("/api/selectrole", middleware.JWTAuthMiddleware(Service.SelectRole))
	e.POST("/api/class/create", middleware.JWTAuthMiddleware(Service.CreateClass))
	e.GET("/api/class/get", middleware.JWTAuthMiddleware(Service.GetClasses))
	e.GET("/api/class/:id", middleware.JWTAuthMiddleware(Service.CheckParticularClass))
	e.POST("/api/class/join", middleware.JWTAuthMiddleware(Service.JoinClass))
	e.POST("/api/exam/create", middleware.JWTAuthMiddleware(Service.SetExam))
	e.GET("/api/exam/getAll", middleware.JWTAuthMiddleware(Service.GetAllExams))
	e.POST("/api/exam/question/add", middleware.JWTAuthMiddleware(Service.AddQuestion))
	e.GET("/api/problems/get", middleware.JWTAuthMiddleware(Service.GetAllProblems))
	e.POST("/api/problems/edit", middleware.JWTAuthMiddleware(Service.RemoveOrAddProblem))
	e.GET("/api/exam/get", middleware.JWTAuthMiddleware(Service.MyExams))
}
