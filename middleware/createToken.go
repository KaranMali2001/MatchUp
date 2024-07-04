package middleware

import (
	
	"log"
	

	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

var JWTKey = []byte("my_secrete_key")

func CreateToken(c echo.Context, role string, username string) (string, error) {
	claims := models.JWTClaims{
		Role:             role,
		Username:         username,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		log.Println(err)
		return "", err
	}
	
	
	return tokenString,nil
}
