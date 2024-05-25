package tournament

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/Database"
	"github.com/labstack/echo"
)

func UpdateTournament(c echo.Context) error {
	db:=database.Db
	var tournament database.Tournament
	name:=c.Param("tournament_name")
	if err:=c.Bind(&tournament);err!=nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,"error while binding")
	}
	updateTournament:=tournament
	if err:=db.Model(&updateTournament).Where("tournament_name = ?",name).Updates(&tournament).Error;err!=nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,"error while updating")
	}
	return c.JSON(http.StatusOK, "updated tournament sucessfully")
}
