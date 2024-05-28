package match

import (
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
)

func UpdateMatch(c echo.Context) error {
	id := c.Param("id")

	var match models.Match
	db := database.Db
	if err := c.Bind(&match); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while binding")
	}
	currMatch := match
	err := db.Model(&currMatch).Where("id = ?", id).Updates(&match).Error
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while updating match")
	}
	return c.JSON(http.StatusOK, "match result updated successfully")
}
