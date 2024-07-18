package tournament

import (
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
)

func GetRegistration(c echo.Context) error {
	tournamentName := c.Param("tournament_name")
	db := database.Db
	var players []string
	result := make(chan error, 1)

	go func() {
		defer close(result)
		if err := db.Model(&models.Registration{}).Where("tournament_name = ?", tournamentName).
			Pluck("player_username", &players).Error; err != nil {
			log.Println(err)
			result <- err
			return
		}
		result <- nil

	}()

	err := <-result
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error while getting reg")
	}
	return c.JSON(http.StatusOK, players)

}
