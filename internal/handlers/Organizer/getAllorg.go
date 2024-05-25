package organizer

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
)

func GetAllOrg(c echo.Context) error {
	db := database.Db
	var org []models.Organizer
	result := db.Find(&org)
	if result.Error != nil {
		log.Println(result.Error)
		return c.JSON(http.StatusInternalServerError, "error while finding")
	}
	return c.JSON(http.StatusFound, org)
}
