package tournament

import (
	"log"
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func GetTournament(c echo.Context) error {
	db := database.Db
	var tournament models.Tournament
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return helper.GetInfo(c, db, &tournament, id)
}
