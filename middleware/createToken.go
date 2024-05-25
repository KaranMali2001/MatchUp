package middleware

import (
	"log"

	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secrate_key")

func CreateToken(role string, username string) (string, error) {
	claims := models.JWTClaims{
		Role:             role,
		Username:         username,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}
