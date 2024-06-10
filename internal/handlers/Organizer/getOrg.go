package organizer

import (
	"log"
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func GetOrganizerInfo(c echo.Context) error {
	var organizer models.Organizer
	db := database.Db
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return helper.GetInfo(c, db, &organizer, id)
}
