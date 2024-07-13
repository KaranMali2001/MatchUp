package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

func VerifyJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("set-cookie")
		if err != nil {
			if err == http.ErrNoCookie {
				log.Println(err)
				return c.JSON(http.StatusBadRequest, "no cookie has been set")
			}
			log.Println(err)
			return c.JSON(http.StatusBadRequest, "error while retriving cookie")

		}
		authHeader := cookie.Value
		fmt.Println(authHeader)
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing jwt")
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
