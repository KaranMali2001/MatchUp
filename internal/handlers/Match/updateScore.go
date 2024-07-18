package match

import (
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func UpdateScore(c echo.Context) error {
	id := c.Param("id")

	var match models.Match
	db := database.Db
	idint, _ := strconv.Atoi(id)
	return helper.UpdateInfo(c, db, &match, idint)

}
