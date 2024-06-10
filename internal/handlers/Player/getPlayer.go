package player

import (
	"log"
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func GetPlayerInfo(c echo.Context) error {
	var player models.Player
	db := database.Db
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return helper.GetInfo(c, db, &player, id)

}
