package organizer

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
)

// soft delete entry still exist in db
func DeleteOranizer(c echo.Context) error {
	db := database.Db
	var organizer models.Organizer
	username := c.Param("username")
	result := db.Where("username = ?", username).Delete(&organizer)
	if result.Error != nil {
		log.Println(result.Error)
		return c.JSON(http.StatusInternalServerError, "error while deleting")
	}
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusOK, "org deleted sucessfully")
	} else {
		return c.JSON(http.StatusNotFound, "org does not exist")
	}

}
