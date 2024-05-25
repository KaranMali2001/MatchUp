package tournament

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"

	"github.com/labstack/echo"
)

func NewTournament(c echo.Context) error {
	db := database.Db
	tournament := new(models.Tournament)
	if err := c.Bind(&tournament); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "binding failed")
	}

	result := db.Create(tournament)
	if result.Error != nil {
		log.Println(result.Error)
		return c.JSON(http.StatusInternalServerError, "error while adding to db")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "new tournament created sucessfully",
		"tournament": tournament,
	})
}
