package tournament

import (
	database "github.com/KaranMali2001/MatchUp/Database"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func AllTournament(c echo.Context) error {
	db := database.Db
	var tournaments []database.Tournament
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
