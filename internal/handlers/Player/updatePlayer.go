package player

import (
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"

	"github.com/labstack/echo"
)

func UpdatePlayer(c echo.Context) error {
	db := database.Db
	var player models.Player
	id := c.Param("id")

	idint, _ := strconv.Atoi(id)
	return helper.UpdateInfo(c, db, &player, idint)

}
