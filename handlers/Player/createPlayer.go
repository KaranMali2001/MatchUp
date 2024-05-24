package player

import (
	"log"
	"net/http"

	database "github.com/KaranMali2001/MatchUp/Database"
	"github.com/labstack/echo"
)

func CreatePlayer(c echo.Context) error {
    db:= database.Db
	player:= new(database.Player)
	if err:=c.Bind(&player);err!=nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,"error while binding the body")
	}
	result:= db.Create(player)
	if err:= result.Error;err!=nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,"error while adding to db")
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"message":"player created sucessfullly",
		"player":player,
	})
}
