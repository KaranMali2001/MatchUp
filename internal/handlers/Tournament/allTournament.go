package tournament

import (
	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func AllTournament(c echo.Context) error {
	db := database.Db
	var tournaments []models.Tournament

	return helper.GetAllInfo(c, db, &tournaments)

}
