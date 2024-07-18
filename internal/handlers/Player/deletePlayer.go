package player

import (
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

// soft delete entry still exist in db
func DeletePlayer(c echo.Context) error {
	var player models.Player
	db := database.Db
	id := c.Param("id")
	idint, _ := strconv.Atoi(id)
	return helper.DeleteInfo(c, db, &player, idint)
}
