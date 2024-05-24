// route should be procted for admin only
package player

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/Database"
	"github.com/labstack/echo"
)

func GetAllPlayer(c echo.Context) error {
	db := database.Db
	var player []database.Player
	result := db.Find(&player)
	if result.Error != nil {
		log.Println(result.Error)
		return c.JSON(http.StatusInternalServerError, "error while finding users")
	}
	return c.JSON(http.StatusFound, player)
}
