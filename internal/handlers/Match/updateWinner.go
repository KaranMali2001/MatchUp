package match

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func UpdateWinner(c echo.Context) error {
	idstr := c.Param("id")
	id, _:= strconv.Atoi(idstr)
	var match models.Match

	db := database.Db
	err := helper.UpdateInfo(c, db, &match, id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while updating")

	}
	// check if all the matches are completed or not
	go func() {
		str := isMatchCompleted(id)
		if str == "success" {
			//go prepare the next round
		} else {
			fmt.Println(str)
		}
	}()
	return c.JSON(http.StatusOK, "match has been updated successfully")
}
