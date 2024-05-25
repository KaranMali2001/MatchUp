package organizer

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/Database"
	"github.com/labstack/echo"
)

func NewOrganizer(c echo.Context) error {
	organizer := new(database.Organizer)
	db := database.Db

	if err := c.Bind(&organizer); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while binding")
	}
	organizer.Role = "organizer"
	result := db.Create(&organizer)
	if result.Error != nil {
		log.Println(result.Error)
		return c.JSON(http.StatusInternalServerError, "error wihle createing new org")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "new org created sucessfully",
		"organizer": organizer,
	})
}
