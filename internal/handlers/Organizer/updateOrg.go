package organizer

import (
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func UpdateOrganizer(c echo.Context) error {
	var organizer models.Organizer
	db := database.Db
	id := c.Param("id")

	idint, _ := strconv.Atoi(id)
	return helper.UpdateInfo(c, db, &organizer, idint)
}
