package organizer

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetOrganizerInfo(c echo.Context) error {
	var organizer models.Organizer
	db := database.Db
	username := c.Param("username")
	err := db.Model(&organizer).Where("username = ?", username).First(&organizer).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println(err)
			return c.JSON(http.StatusNotFound, "record not found")
		}
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while finding user")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "this is organization info",
		"organizer": organizer,
	})
}
