package tournament

import (
	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
	"net/http"
)

func UpdateTournament(c echo.Context) error {
	db := database.Db
	var tournament models.Tournament
	claims := c.Get("claims").(*models.JWTClaims)
	if claims.Role != "organizer" {
		return c.JSON(http.StatusUnauthorized, "only org can create tournament")
	}
	id := c.Param("id")
	return helper.UpdateInfo(c, db, &tournament, id)

}
