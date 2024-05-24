package player

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/Database"
	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func UpdatePlayer(c echo.Context) error {
	db := database.Db
	var player database.Player
	username := c.Param("username")

	if err := c.Bind(&player); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while binding")
	}
	//write validation for email explicitly
	updatePlayer := player
	if err := db.Model(&updatePlayer).Where("username = ?", username).Updates(&player).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while updating")
	}
	return c.JSON(http.StatusOK, player)
}
