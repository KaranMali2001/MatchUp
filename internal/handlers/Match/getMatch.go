package match

import (
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
)

func GetMatch(c echo.Context) error {
	id := c.Param("id")
	db := database.Db
	var match models.Match
	err := db.Where("id = ?", id).First(&match).Error
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while finding ")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "match info",
		"match":   match,
	})
}
