package routes

import (
	"immersive/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	loginhandler "immersive/domains/login/handlers"
	loginrepo "immersive/domains/login/repositories"
	loginusecase "immersive/domains/login/usecases"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, cfg *config.AppConfig) {
	/*
		Dependency Injection
	*/

	loginRepo := loginrepo.New(db)
	loginUsecase := loginusecase.New(loginRepo)
	loginHandler := loginhandler.New(loginUsecase)

	/*
		Routes
	*/
	e.POST("/login", loginHandler.Login)
}
