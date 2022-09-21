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

	/*
		Routes
	*/
	e.POST("/login", loginHandler.Login)

	e.POST("/mentor", mentorHandler.Create, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.GET("/mentors", mentorHandler.ReadAll, middlewares.JWTMiddleware())
	e.GET("/mentor/:id", mentorHandler.ReadById, middlewares.JWTMiddleware())
	e.PUT("/mentor/:id", mentorHandler.Update, middlewares.JWTMiddleware())
	e.DELETE("/mentor/:id", mentorHandler.Delete, middlewares.JWTMiddleware(), middlewares.IsAdmin)
}
