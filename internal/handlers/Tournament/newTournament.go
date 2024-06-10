package tournament

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"

	"github.com/labstack/echo"
)

func NewTournament(c echo.Context) error {
	db := database.Db
	tournament := new(models.Tournament)
	claims := c.Get("claims").(*models.JWTClaims)
	//validate
	if err := c.Bind(&tournament); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "binding failed")
	}
	tournament.OrganizerName = claims.Username
	if claims.Role != "organizer" {
		return c.JSON(http.StatusUnauthorized, "only org can create tournament")
	}
	errChannel := helper.Create(db, tournament)
	// Wait for the result from the channel
	if err := <-errChannel; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error while creating tournament")
	}

	return c.JSON(http.StatusOK, "Tournament created successfully")

}
