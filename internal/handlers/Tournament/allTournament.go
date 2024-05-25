package tournament

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
)

func AllTournament(c echo.Context) error {
	db := database.Db
	var tournaments []models.Tournament
	result := db.Find(&tournaments)
	if result.Error != nil {
		log.Println(result.Error)
		return c.JSON(http.StatusInternalServerError, "error while finding users")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "returning all the tournaments",
		"tournaments": tournaments,
	})
}
