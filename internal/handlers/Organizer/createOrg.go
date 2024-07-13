package organizer

import (
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"

	"github.com/labstack/echo"
)

func NewOrganizer(c echo.Context) error {
	organizer := new(models.Organizer)
	db := database.Db

	if err := c.Bind(&organizer); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while binding")
	}
	organizer.Role = "organizer"

	errChannel := helper.Create(db, organizer)
	// Wait for the result from the channel
	if err := <-errChannel; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error while creating tournament")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "new org created successfully",
		"organizer": organizer,
	})
}
