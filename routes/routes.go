package routes

import (
	"immersive/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	loginhandler "immersive/domains/login/handlers"
	loginrepo "immersive/domains/login/repositories"
	loginusecase "immersive/domains/login/usecases"

	mentorhandler "immersive/domains/mentor/handlers"
	mentorrepo "immersive/domains/mentor/repositories"
	mentorusecase "immersive/domains/mentor/usecases"

	classhandler "immersive/domains/class/handlers"
	classrepo "immersive/domains/class/repositories"
	classusecase "immersive/domains/class/usecases"

	menteehandler "immersive/domains/mentee/handlers"
	menteerepo "immersive/domains/mentee/repositories"
	menteeusecase "immersive/domains/mentee/usecases"

	"immersive/middlewares"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, cfg *config.AppConfig) {
	/*
		Dependency Injection
	*/

	loginRepo := loginrepo.New(db)
	loginUsecase := loginusecase.New(loginRepo)
	loginHandler := loginhandler.New(loginUsecase)

	mentorRepo := mentorrepo.New(db)
	mentorUsecase := mentorusecase.New(mentorRepo)
	mentorHandler := mentorhandler.New(mentorUsecase)

	classRepo := classrepo.New(db)
	classUsecase := classusecase.New(classRepo)
	classHandler := classhandler.New(classUsecase)

	menteeRepo := menteerepo.New(db)
	menteeUsecase := menteeusecase.New(menteeRepo)
	menteeHandler := menteehandler.New(menteeUsecase)

	/*
		Routes
	*/
	e.POST("/login", loginHandler.Login)

	e.POST("/mentor", mentorHandler.Create, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.GET("/mentors", mentorHandler.ReadAll, middlewares.JWTMiddleware())
	e.GET("/mentor/:id", mentorHandler.ReadById, middlewares.JWTMiddleware())
	e.PUT("/mentor/:id", mentorHandler.Update, middlewares.JWTMiddleware())
	e.DELETE("/mentor/:id", mentorHandler.Delete, middlewares.JWTMiddleware(), middlewares.IsAdmin)

	e.POST("/class", classHandler.Create, middlewares.JWTMiddleware())
	e.GET("/classes", classHandler.GetAll, middlewares.JWTMiddleware())
	e.GET("/class/:id", classHandler.Get, middlewares.JWTMiddleware())
	e.PUT("/class/:id", classHandler.Update, middlewares.JWTMiddleware())
	e.DELETE("/class/:id", classHandler.Delete, middlewares.JWTMiddleware())

	e.POST("/mentee", menteeHandler.Create, middlewares.JWTMiddleware())
	e.GET("/mentees", menteeHandler.GetAll, middlewares.JWTMiddleware())
	e.GET("/mentee/:id", menteeHandler.Get, middlewares.JWTMiddleware())
	e.PUT("/mentee/:id", menteeHandler.Update, middlewares.JWTMiddleware())
	e.DELETE("/mentee/:id", menteeHandler.Delete, middlewares.JWTMiddleware())
}
