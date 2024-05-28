package match

import (
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
)

func GetAllMatches(c echo.Context) error {
	tournamentName := c.Param("tournament_name")
	var matches []models.Match
	db := database.Db
	err := db.Where("tournament_name = ? ", tournamentName).Find(&matches).Error
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while finding matches")
	}
	return c.JSON(http.StatusOK, matches)
}
