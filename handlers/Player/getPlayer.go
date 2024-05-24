package player

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/Database"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetPlayerInfo(c echo.Context) error {
	var player database.Player
	db := database.Db
	username := c.Param("username")
	err := db.Model(&player).Where("username = ?", username).First(&player).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println(err)
			return c.JSON(http.StatusNotFound, "error not found")
		}
		return c.JSON(http.StatusInternalServerError, "error while finding the data")
	}
	return c.JSON(http.StatusOK, player)
}
