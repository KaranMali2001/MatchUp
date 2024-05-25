package tournament

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/Database"
	"gorm.io/gorm"

	"github.com/labstack/echo"
)

func NewTournament(c echo.Context) error {
	db := database.Db
	tournament := new(database.Tournament)
	if err := c.Bind(&tournament); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "binding failed")
	}
	var org database.Organizer

	err := db.Where("organizer_name = ? ", tournament.OrganizerName).First(&org).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "org does not exist")
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
