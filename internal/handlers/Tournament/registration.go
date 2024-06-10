package tournament

import (
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func Registration(c echo.Context) error {
	reg := new(models.Registration)
	db := database.Db
	var tournament models.Tournament

	TournamentName := c.Param("tournament_name")
	claims := c.Get("claims").(*models.JWTClaims)
	if claims.Role != "player" {
		return c.JSON(http.StatusUnauthorized, "you are not player")
	}
	if err := db.Model(&tournament).Where("live = ?", true).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "tournament is not live")
	}
	reg.TournamentName = TournamentName
	reg.PlayerUsername = claims.Username
	errChannel := helper.Create(db, reg)
	if err := <-errChannel; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while creating")
	}
	tournament.TotalPlayer += 1
	return c.JSON(http.StatusOK, "registration for tournament is successful")
}
