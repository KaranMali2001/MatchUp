package homepage

import (
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/middleware"
	"github.com/golang-jwt/jwt/v4"

	"github.com/labstack/echo"
)

// use react and give option as to sign up as player or org and then redirect the root
func Login(c echo.Context) error {
	db := database.Db
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user struct {
		Username string
		Password string
		Role     string
	}
	if err := db.Table("players").Select("username, password, role").Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		if err := db.Table("organizers").Select("username, password, role").Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, "user name password not found in db")
		}
	}
	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JWTKey)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while creating jwt")
	}
	return c.JSON(http.StatusOK, tokenString)
}
