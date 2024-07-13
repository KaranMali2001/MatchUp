package homepage

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/middleware"

	"github.com/labstack/echo"
)

// use react and give option as to sign up as player or org and then redirect the root
func Login(c echo.Context) error {
	db := database.Db

	var user struct {
		Username string
		Password string
		Role     string
	}
	if err := c.Bind(&user); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "not a good request")
	}
	fmt.Println("username is :", user.Username)
	fmt.Println(user.Password)
	if err := db.Table("players").Select("username, password, role").Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; err != nil {
		if err := db.Table("organizers").Select("username, password, role").Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; err != nil {
			log.Println(err)
			return c.JSON(http.StatusUnauthorized, "user name password not found in db")
		}
	}
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	err := middleware.CreateToken(c, user.Role, user.Username)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "cant log in at this time")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "logged in successfully",
	})
}
