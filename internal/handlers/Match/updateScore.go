package match

import (
	"net/http"
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"

	//"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func UpdateScore(c echo.Context) error {
    id := c.Param("id")
    idint, err := strconv.Atoi(id)
    db:=database.Db
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid match ID"})
    }


    return helper.UpdateInfo(c,db,&models.Match{},idint)
}
