package tournament

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetTournament(c echo.Context) error {
	db := database.Db
	var tournament models.Tournament
	name := c.Param("tournament_name")
	err := db.Model(&tournament).Where("tournament_name = ?", name).First(&tournament).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println(err)
			return c.JSON(http.StatusNotFound, "error not found")
		}
		return c.JSON(http.StatusInternalServerError, "error while finding the data")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"tournament": tournament,
	})
}
