package player

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
)

// soft delete entry still exist in db
func DeletePlayer(c echo.Context) error {
	var player models.Player
	db := database.Db
	username := c.Param("username")
	result := db.Where("username = ?", username).Delete(&player)
	if result.Error != nil {
		log.Println(result.Error)
		return c.JSON(http.StatusInternalServerError, "error while deleting player")
	}
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusOK, "player deleted successfully")
	} else {
		return c.JSON(http.StatusNotFound, "player does not exist")
	}
}
