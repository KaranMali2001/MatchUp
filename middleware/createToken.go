package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

var JWTKey = []byte("my_secrete_key")

func CreateToken(c echo.Context, role string, username string, id string) error {
	claims := models.JWTClaims{
		Role:             role,
		Username:         username,
		ID:               id,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		log.Println(err)
		return err
	}
	cookie := http.Cookie{
		Name:    "set-cookie",
		Value:   tokenString,
		Expires: time.Now().Add(24 * time.Hour),

		Path: "/",
	}

	c.SetCookie(&cookie)

	return nil
}
