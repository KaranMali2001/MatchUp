package tournament

import (
	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func DeleteTournament(c echo.Context) error {
	db := database.Db
	var tournament models.Tournament
	id := c.Param("id")
	return helper.DeleteInfo(c, db, &tournament, id)

}
