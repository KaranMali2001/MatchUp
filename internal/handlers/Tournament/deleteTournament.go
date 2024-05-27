package tournament

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
)

func DeleteTournament(c echo.Context) error {
	db := database.Db
	var tournament models.Tournament
	name := c.Param("tournament_name")
	result := db.Where("tournament_name = ?", name).Delete(&tournament)
	if result.Error != nil {
		log.Println(result.Error)
		return c.JSON(http.StatusInternalServerError, "error while deleting")
	}
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusOK, "deleted successfully")
	} else {
		return c.JSON(http.StatusNotFound, "tournament not found")
	}

}
