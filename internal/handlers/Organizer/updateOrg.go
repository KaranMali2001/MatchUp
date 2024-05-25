package organizer

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
)

func UpdateOrganizer(c echo.Context) error {
	var organizer models.Organizer
	db := database.Db
	username := c.Param("username")
	if err := c.Bind(&organizer); err != nil {
		return c.JSON(http.StatusInternalServerError, "error while binding")
	}
	updateOrg := organizer
	err := db.Model(&updateOrg).Where("username = ?", username).Updates(&organizer).Error
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while updating")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "org info updated sucessfully",
		"organizer": organizer,
	})
}
