package match

import (
	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func UpdateMatch(c echo.Context) error {
	id := c.Param("id")

	var match models.Match
	db := database.Db
	return helper.UpdateInfo(c, db, &match, id)

}
