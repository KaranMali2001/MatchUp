package organizer

import (
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

// soft delete entry still exist in db
func DeleteOrganizer(c echo.Context) error {
	db := database.Db
	var organizer models.Organizer
	id := c.Param("id")
	idint, _ := strconv.Atoi(id)
	return helper.DeleteInfo(c, db, &organizer, idint)

}
