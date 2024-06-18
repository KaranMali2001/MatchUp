package match

import (
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func UpdateWinner(c echo.Context) error {
	id := c.Param("id")

	var match models.Match
	
	db := database.Db
	err:=helper.UpdateInfo(c,db,&match,id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,"error while updating")

	}
	go func ()  {
		Rounds(db,id)
	}()
return c.JSON(http.StatusOK,"match has been updated successfully")
}
