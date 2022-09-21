package middlewares

import (
	"errors"
	"immersive/config"
	"immersive/utils/helpers"
	"net/http"
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

func CreateToken(userID uint, role string) (string, error) {
	if userID == 0 || role == "" {
		return "", errors.New("empty response")
	}

	cfg := config.GetConfig()
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT_SECRET))
}

func ExtractToken(c echo.Context) (int, string) {
	token, ok := c.Get("user").(*jwt.Token)

	if !ok {
		c.JSON(http.StatusForbidden, helpers.FailedResponse("jwt not ok"))
	}

	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		uid := claim["userID"].(float64)
		role := claim["role"].(string)
		return int(uid), role
	}

	return 0, ""
}
