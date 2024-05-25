package organizer

import (
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/middleware"
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
	result := db.Create(&organizer)
	if result.Error != nil {
		log.Println(result.Error)
		return c.JSON(http.StatusInternalServerError, "error wihle createing new org")
	}
	token, err := middleware.CreateToken("organizer", organizer.Username)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "entry created but error while creating jwt token")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "new org created sucessfully",
		"organizer": organizer,
		"token":     token,
	})
}