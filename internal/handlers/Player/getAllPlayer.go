package player

import (
	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func GetAllPlayer(c echo.Context) error {
	db := database.Db
	var player []models.Player
	return helper.GetAllInfo(c, db, &player)
}
