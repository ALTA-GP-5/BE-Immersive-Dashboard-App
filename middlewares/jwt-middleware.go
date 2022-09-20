package middlewares

import (
	"immersive/config"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	cfg := config.GetConfig()
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(cfg.JWT_SECRET),
	})
}

func CreateToken(userID uint) (string, error) {
	cfg := config.GetConfig()
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT_SECRET))
}

func GetToken(c echo.Context) (int, error) {
	token, ok := c.Get("user").(*jwt.Token)

	if !ok {
		log.Fatal(ok)
	}

	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		uid := claim["userID"].(float64)
		return int(uid), nil
	}

	return 0, fmt.Errorf("token invalid")
}
