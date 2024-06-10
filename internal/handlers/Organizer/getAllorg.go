package organizer

import (
	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func GetAllOrg(c echo.Context) error {
	db := database.Db
	var org []models.Organizer
	return helper.GetAllInfo(c, db, &org)
}
