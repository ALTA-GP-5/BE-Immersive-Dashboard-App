package main

import (
	"immersive/config"
	"immersive/exceptions"

	database "immersive/utils/database/gorm-mysql"

	"immersive/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidation struct {
	validator *validator.Validate
}

func (cv *CustomValidation) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidation{validator: validator.New()}
	e.HTTPErrorHandler = exceptions.CustomErrorHandling
	e.Use(middleware.Recover())

	cfg := config.GetConfig()
	db := database.InitDB(cfg)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	routes.InitRoutes(e, db, cfg)

	err := e.Start(":" + cfg.SERVER_PORT)

	if err != nil {
		panic(err)
	}
}
