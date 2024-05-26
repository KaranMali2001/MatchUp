package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

func VerifyJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing jwt")
		}

		if strings.HasPrefix(authHeader, "Bearer ") {
			authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			return c.JSON(http.StatusUnauthorized, "does not contain bearer")
		}
		token, err := jwt.ParseWithClaims(authHeader, &models.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(JWTKey), nil

		})
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "error while validating token")
		}
		claims := token.Claims.(*models.JWTClaims)
		c.Set("claims", claims)
		return next(c)
	}
}
