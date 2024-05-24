package organizer

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/Database"
	"github.com/labstack/echo"
)

func GetAllOrg(c echo.Context) error {
	db := database.Db
	var org []database.Organizer
	result := db.Find(&org)
	if result.Error != nil {
		log.Println(result.Error)
		return c.JSON(http.StatusInternalServerError, "error while finding")
	}
	return c.JSON(http.StatusFound, org)
}
